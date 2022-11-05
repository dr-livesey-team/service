import React from "react";
import { YMaps, Map, Clusterer, Placemark } from "react-yandex-maps";

import POINTS from "./data";

const mapState = {
  center: [55.751574, 37.573856],
  zoom: 10,
  controls: [],
};

export class CustomMap extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      selectedPoint: null,
      iconLayoutTemplate: null,
      map: null,
    };
  }

  clusterIconShape = {
    type: "Circle",
    coordinates: [0, 0],
    radius: 19,
  };

  onPlacemarkClick = (point) => () => {
    this.setState({ selectedPoint: point });
  };

  render() {
    const { selectedPoint } = this.state;

    return (
      <div className="App">
        <YMaps query={{ lang: "ru_RU", load: "package.full" }}>
          <Map
            height={"600px"}
            width={"100%"}
            defaultState={mapState}
            modules={[
              "templateLayoutFactory",
              "option.presetStorage",
              "option.Manager",
              "control.ZoomControl",
              "control.FullscreenControl",
            ]}
            instanceRef={(ref) => (this.map = ref)}
            onLoad={(ymaps) =>
              this.setState({
                iconLayoutTemplate:
                  ymaps.templateLayoutFactory.createClass("<div></div>"),
              })
            }
          >
            <Clusterer
              options={{
                hasBalloon: true,
                hasHint: false,
              }}
            >
              {POINTS.map((point, index) => (
                <Placemark
                  key={index}
                  geometry={[point[0], point[1]]}
                  onClick={() => {
                      this.onPlacemarkClick(point);
                      this.props.setActivePoint(point);
                  }}
                  properties={{
                    item: index,
                  }}
                  options={{
                    iconLayout: "default#image",
                    iconImageSize: [35, 35],
                    iconImageHref: "https://ildan-dev.ru/placemark.svg",
                  }}
                />
              ))}
            </Clusterer>
          </Map>
        </YMaps>
        {selectedPoint && (
          <div>
            <h1>Selected point: {selectedPoint.title}</h1>
            <p>{selectedPoint.descr}</p>
          </div>
        )}
      </div>
    );
  }
}

export default CustomMap;
