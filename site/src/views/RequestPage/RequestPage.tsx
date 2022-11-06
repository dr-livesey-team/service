import React, {useEffect, useState} from "react";
import S from "./requestPage.module.scss";
import Comment from "./Comment/Comment";
import axios from "axios";
import { infoMock, serverURL } from "../../assets/requestMock";

const RequestPage: React.FC<{ id: number }> = ({ id }) => {
  const [requestInfo, setRequestInfo] = useState<any>([]);
  useEffect(() => {
    axios
      .get(serverURL + "info?id=" + id)
      .then((response) => setRequestInfo(response.data?.requests))
      .catch(() => setRequestInfo(infoMock));
  }, []);
  return (
    <div className={S.page__wrapper}>
      <div className={S.header}>
        <div className={S.close}>Назад</div>
        <h2 className={S.title}>Отсутствие отопления в комнате, квартире</h2>
      </div>
      <div className={S.content}>
        <div className={S.info}>
          <div className={S.info__top}>
            <div className={S.generalInfo}>
              <h2>Описание</h2>
              <h3>Общее время разрешения ситуации</h3>
              <p>2д 4ч</p>
              <h3>Адрес</h3>
              <p>г. Москва, ул. Мира, 5, кв. 4</p>
              <h3>Управляющая организация</h3>
              <p>ГБУ “Жилищник района Теплый стан”</p>
              <h3>Обслуживающая организация</h3>
              <p>ПАО "МОЭК"</p>
            </div>
            <div className={S.abnormalBlock}>
              <h2>Аномалии</h2>
              <ul>
                <li>Не выполнена с первого раза</li>
                <li>Негативный комментарий</li>
                <li>Долгий срок закрытия</li>
                <li>Долгое закрытие</li>
              </ul>
            </div>
          </div>
          <div className={S.info__bottom}>
            <h2>История</h2>
            <table>
              <tr>
                <th>Дата создания</th>
                <th>Время выполнения</th>
                <th>Результативность</th>
                <th>Вид работ</th>
                <th>Описание</th>
                <th></th>
              </tr>
              <Comment />
              <Comment />
              <Comment />
              <Comment />
            </table>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RequestPage;
