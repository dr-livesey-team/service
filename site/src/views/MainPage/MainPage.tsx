import React, { useEffect, useState } from "react";
import S from "./mainPage.module.scss";
import RequestBlock from "./RequestBlock/RequestBlock";
import SearchInput from "../../components/SearchInput/SearchInput";
import FilterModal from "../../components/FilterModal/FilterModal";
import CustomCheckbox from "../../components/CustomCheckbox/CustomCheckbox";
import MapModal from "../../components/MapModal/MapModal";
import axios from "axios";
import { anomaliesMock, serverURL } from "../../assets/requestMock";

const MainPage = () => {
  const [filterActive, setFilterActive] = useState(false);
  const [mapActive, setMapActive] = useState(false);
  const [anomalies, setAnomalies] = useState<Array<any>>([]);
  const [actualSearch, setActualSearch] = useState('');

  useEffect(() => {
    axios
      .get(serverURL + "anomalies")
      .then((response) => setAnomalies(response.data?.anomalies))
      .catch(() => setAnomalies(anomaliesMock.anomalies));
  }, []);

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
          <p className={S.title__count}>3000 шт</p>
        </div>
        <div className={S.header__wrapper}>
          <SearchInput actualSearch={actualSearch} setActualSearch={setActualSearch}/>
          <div className={S.sort} onClick={() => setFilterActive(true)}></div>
        </div>
      </div>
      <div className={S.content}>
        <div className={S.content__map}>
          <div className={S.mapButton} onClick={() => setMapActive(true)}>
            Найти на карте
          </div>
        </div>
        <RequestBlock anomalies={anomalies} />
      </div>
      {filterActive && <FilterModal setActive={setFilterActive} />}
      {mapActive && (
        <MapModal anomalies={anomalies} setActive={setMapActive} setFilterActive={setFilterActive} actualSearch={actualSearch} setActualSearch={setActualSearch}/>
      )}
    </div>
  );
};

export default MainPage;
