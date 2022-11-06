import React, { useEffect, useState } from "react";
import S from "./mapModal.module.scss";
import SearchInput from "../SearchInput/SearchInput";
import CustomCheckbox from "../CustomCheckbox/CustomCheckbox";
import CustomMap from "../CustomMap/CustomMap";
import Request from "../../views/MainPage/RequestBlock/Request/Request";

const MapModal: React.FC<{
  setActive: any;
  setFilterActive: any;
  anomalies: any;
}> = ({ setActive, setFilterActive, anomalies }) => {
  const [activePoint, setActivePoint] = useState(undefined);

  useEffect(() => {
    console.log(activePoint);
  }, [activePoint]);

  return (
    <div className={S.modal__wrapper}>
      <div className={S.close} onClick={() => setActive(false)}>
        К списку заявок
      </div>
      <div className={S.modal}>
        <div className={S.header}>
          <div className={S.title}>
            <h2 className={S.title__header}>Поиск по карте</h2>
          </div>
          <div className={S.header__wrapper}>
            <SearchInput />
            <div className={S.sort} onClick={() => setFilterActive(true)}></div>
            <CustomCheckbox>Только аномальные</CustomCheckbox>
          </div>
        </div>
        <CustomMap anomalies={anomalies} setActivePoint={setActivePoint} />
        <div className={S.requestList}>
          <table className={S.requestBlock}>
            <th className={S.header}>
              <td>Аномальная</td>
              <td>Наименование</td>
              <td>Дата создания</td>
            </th>
            {anomalies.anomalies.map(() => {
              <Request minRequest={true} />
            })}
          </table>
        </div>
      </div>
    </div>
  );
};

export default MapModal;
