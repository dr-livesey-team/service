import React, { useEffect, useState } from "react";
import S from "./requestPage.module.scss";
import Comment from "./Comment/Comment";
import axios from "axios";
import { infoMock, serverURL } from "../../assets/requestMock";
import { Link } from "react-router-dom";
import { anomaly_category_list } from "../../assets/filterLists";

const RequestPage: React.FC = () => {
  const [requestInfo, setRequestInfo] = useState<Array<any>>([]);
  const [date, setDate] = useState<Array<number>>([]);
  const [buf, setBuf] = useState(0);
  useEffect(() => {
    let loc = document.location.search;
    loc.slice(0, 3);
    axios
      .get(serverURL + "info" + loc)
      .then((response) => {
        setRequestInfo(response.data?.requests);
        setBuf(
          new Date(
            parseInt(requestInfo[0]?.closing_date.slice(0, 4)),
            (parseInt(requestInfo[0]?.closing_date.slice(5, 7)) + 1) % 12,
            parseInt(requestInfo[0]?.closing_date.slice(8, 10)),
            parseInt(requestInfo[0]?.closing_date.slice(11, 13)),
            parseInt(requestInfo[0]?.closing_date.slice(14, 16))
          ).getTime() -
            new Date(
              parseInt(requestInfo[0]?.opening_date.slice(0, 4)),
              (parseInt(requestInfo[0]?.opening_date.slice(5, 7)) + 1) % 12,
              parseInt(requestInfo[0]?.opening_date.slice(8, 10)),
              parseInt(requestInfo[0]?.opening_date.slice(11, 13)),
              parseInt(requestInfo[0]?.opening_date.slice(14, 16))
            ).getTime()
        );
      })
      .catch(() => setRequestInfo([infoMock]));
  }, []);
  useEffect(() => {
    setDate([
      Math.floor(buf / 1000 / 60 / 60 / 24),
      Math.floor((buf / 1000 / 60 / 60) % 24),
      Math.floor((buf / 1000 / 60) % 60),
    ]);
  }, []);
  return (
    <div className={S.page__wrapper}>
      <div className={S.header}>
        <Link to={"/monitoring"} className={S.close}>
          Назад
        </Link>
        <h2 className={S.title}>{requestInfo[0]?.fault_name}</h2>
      </div>
      <div className={S.content}>
        <div className={S.info}>
          <div className={S.info__top}>
            <div className={S.generalInfo}>
              <h2>Описание</h2>
              <h3>Время выполнения</h3>
              <p>
                {date[0] || 0}д {date[1] || 0}ч {date[2] || 0}мин
              </p>
              <h3>Адрес</h3>
              <p>{requestInfo[0]?.address}</p>
              <h3>Управляющая организация</h3>
              <p>{requestInfo[0]?.management_company_name}</p>
              <h3>Обслуживающая организация</h3>
              <p>{requestInfo[0]?.service_organization_name}</p>
            </div>
            <div className={S.abnormalBlock}>
              <h2>Аномалии</h2>
              <ul>
                <li>
                  {anomaly_category_list[requestInfo[0]?.anomaly_category + 1]}
                </li>
              </ul>
            </div>
          </div>
          <div className={S.info__bottom}>
            <h2>История</h2>
            <table>
              <tr>
                <th>Номер заявки</th>
                <th>Дата создания</th>
                <th>Дата закрытия</th>
                <th>Результативность</th>
                <th>Описание</th>
                <th></th>
              </tr>
              {requestInfo.map((item, index) => (
                <Comment item={item} key={index}/>
              ))}
            </table>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RequestPage;
