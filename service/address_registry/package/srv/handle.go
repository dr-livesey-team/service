package srv

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/dr-livesey-team/service/service/address_registry/package/util"
)

const (
	UrlFmt      string = "https://apidata.mos.ru/v1/datasets/60562/rows?$top=1&$filter=Cells/SIMPLE_ADDRESS%%20eq%%20%s&api_key=90dd708a35a983615cdbfaf42515aff9"
	ContentType string = "application/json"
	Body        string = "[\"geoData\"]"
)

func Handle(conn net.Conn) {
	defer conn.Close()

	for {
		buffer, err := ReadRequest(conn)
		if err != nil {
			util.LogError(err)
			return
		}
		util.Log(util.Debug, "Received buffer is '%s'", string(buffer))

		request, err := UnmarshalRequest(buffer)
		if err != nil {
			util.LogError(err)
			return
		}
		util.Log(util.Debug, "Unmarshaled request is {Address: '%s'}", request.Address)

		response, err := Process(request)
		if err != nil {
			util.LogError(err)
			return
		}
		util.Log(util.Debug, "Response to marshal is {Latitude: %g, Longitude: %g}\n", response.Latitude, response.Longitude)

		buffer, err = MarshalResponse(response)
		if err != nil {
			util.LogError(err)
			return
		}
		util.Log(util.Debug, "Buffer to send is '%s'", string(buffer))

		err = WriteResponse(conn, buffer)
		if err != nil {
			util.LogError(err)
			return
		}
	}
}

func ReadRequest(conn net.Conn) ([]byte, error) {
	reader := bufio.NewReader(conn)

	buffer, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func WriteResponse(conn net.Conn, buffer []byte) error {
	writer := bufio.NewWriter(conn)

	_, err := writer.Write(buffer)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return nil
	}

	return nil
}

func Process(request *Request) (*Response, error) {
	url := PrepareRequestUrl(request.Address)
	util.Log(util.Debug, "URL is '%s'\n", url)

	body, err := PerformRequest(url)
	if err != nil {
		util.LogError(err)
		return nil, err
	}
	util.Log(util.Debug, "Body is '%s'\n", body)

	records, err := Unmarshal(body)
	if err != nil {
		util.LogError(err)
		return nil, err
	}
	util.Log(util.Debug, "Number of coordinates is %d\n", len(records))

	latitude := float64(0)
	longitude := float64(0)

	if records[0].Cells.Data.Coordinates != nil {
		latitude = records[0].Cells.Data.Coordinates[0][0][1]
		longitude = records[0].Cells.Data.Coordinates[0][0][0]
	}

	return &Response{Latitude: latitude, Longitude: longitude}, nil
}

func PrepareRequestUrl(address string) string {
	replacer := strings.NewReplacer(" ", "%20")
	return fmt.Sprintf(UrlFmt, replacer.Replace(address))
}

func PerformRequest(url string) ([]byte, error) {
	reader := strings.NewReader(Body)

	response, err := http.Post(url, ContentType, reader)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("apidata.mos.ru error occured")
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

type GeoData struct {
	Coordinates [][][]float64 `json:"coordinates"`
	Shape       string        `json:"type"`
}

type Cells struct {
	Data GeoData `json:"geoData"`
}

type Record struct {
	GlobalId int   `json:"global_id"`
	Number   int   `json:"Number"`
	Cells    Cells `json:"Cells"`
}

func Unmarshal(body []byte) ([]Record, error) {
	records := []Record{}

	err := json.Unmarshal(body, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}
