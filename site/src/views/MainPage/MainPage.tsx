import React, { useEffect, useState } from "react";
import S from "./mainPage.module.scss";
import RequestBlock from "./RequestBlock/RequestBlock";
import SearchInput from "../../components/SearchInput/SearchInput";
import FilterModal from "../../components/FilterModal/FilterModal";
import MapModal from "../../components/MapModal/MapModal";
import axios from "axios";
import { anomaliesMock, serverURL } from "../../assets/requestMock";
import {
  district_name_list,
  management_company_name_list,
  service_organization_name_list, urgency_category_list
} from "../../assets/filterLists";

const MainPage = () => {
  const [filterActive, setFilterActive] = useState(false);
  const [mapActive, setMapActive] = useState(false);
  const [anomalies, setAnomalies] = useState<Array<any>>([]);
  const [actualSearch, setActualSearch] = useState("");
  const [selectedItems, setSelectedItems] = useState<Array<Array<boolean>>>(
    Array(Array(2500), Array(130), Array(2), Array(2900), Array(20))
  );

  useEffect(() => {
    axios
      .get(
        serverURL +
          "anomalies?" +
          selectedItems[0]
            .map((item, index) =>
              !!item
                ? "management_company_name=" +
                  management_company_name_list[index]
                : ""
            )
            .join("") +
          selectedItems[1]
            .map((item, index) =>
              !!item
                ? "&district_name=" +
                  district_name_list[index]
                : ""
            )
            .join("") +
          selectedItems[3]
              .map((item, index) =>
                  !!item
                      ? "&service_organization_name=" +
                      service_organization_name_list[index]
                      : ""
              )
              .join("") +
          selectedItems[4]
              .map((item, index) =>
                  !!item
                      ? "&urgency_category_name=" +
                      urgency_category_list[index]
                      : ""
              )
              .join("")
      )
      .then((response) => setAnomalies(response.data?.anomalies))
      .catch(() => setAnomalies(anomaliesMock.anomalies));
  }, [filterActive]);

  useEffect(() => {
    if (filterActive || mapActive) {
      document.body.style.overflow = "hidden";
      window.scrollTo(0, 0);
    } else document.body.style.overflow = "";
  }, [filterActive, mapActive]);

  return (
    <div className={S.mainPage}>
      <div className={S.header}>
        <div className={S.title}>
          <h2 className={S.title__header}>Заявки</h2>
          <p className={S.title__count}>{anomalies.length} шт</p>
        </div>
        <div className={S.header__wrapper}>
          <SearchInput
            actualSearch={actualSearch}
            setActualSearch={setActualSearch}
          />
          <div className={S.sort} onClick={() => setFilterActive(true)}></div>
        </div>
      </div>
      <div className={S.content}>
        <div className={S.content__map}>
          <div className={S.mapButton} onClick={() => setMapActive(true)}>
            Найти на карте
          </div>
        </div>
        <RequestBlock
          anomalies={anomalies.filter((item) =>
            item?.fault_name.toLowerCase().includes(actualSearch.toLowerCase())
          )}
        />
      </div>
      {filterActive && (
        <FilterModal
          selectedItems={selectedItems}
          setSelectedItems={setSelectedItems}
          setActive={setFilterActive}
        />
      )}
      {mapActive && (
        <MapModal
          anomalies={anomalies}
          setActive={setMapActive}
          setFilterActive={setFilterActive}
          actualSearch={actualSearch}
          setActualSearch={setActualSearch}
        />
      )}
    </div>
  );
};

export default MainPage;
