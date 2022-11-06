import React, { useState } from "react";
import S from "./filterModal.module.scss";
import SearchInput from "../SearchInput/SearchInput";
import FilterItem from "./FilterItem/FilterItem";
import Calendar from "./Calendar/Calendar";

const factorsList = [
  "Адрес дома",
  "Управляющая компания",
  "Район",
  "Дата создания/закрытия заявки",
  "Обслуживающая организация",
  "Категория срочности",
];

const FilterModal: React.FC<{ setActive: any }> = ({ setActive }) => {
  const [activeFactor, setActiveFactor] = useState(0);
  const [actualSearch, setActualSearch] = useState('');

  return (
    <div className={S.modal_wrapper}>
      <div className={S.close} onClick={() => setActive(false)}>
        Закрыть
      </div>
      <div className={S.modal}>
        <div className={S.modal__left}>
          {activeFactor !== 3 ? <SearchInput actualSearch={actualSearch} setActualSearch={setActualSearch}/> : <></>}
          {activeFactor === 3 ? (
            <Calendar />
          ) : (
            <ul className={S.filterList}>
              <FilterItem>Регион</FilterItem>
              <FilterItem>Регион</FilterItem>
              <FilterItem>Регион</FilterItem>
              <FilterItem>Регион</FilterItem>
              <FilterItem>Регион</FilterItem>
              <FilterItem>Регион</FilterItem>
              <FilterItem>Регион</FilterItem>
              <FilterItem>Регион</FilterItem>
              <FilterItem>Регион</FilterItem>
            </ul>
          )}
        </div>
        <div className={S.modal__right}>
          <h3>Отфильтровать по</h3>
          <ul className={S.filterFactors}>
            {factorsList.map((factor, index) => {
              return (
                <li
                  id={index.toString()}
                  onClick={() => setActiveFactor(index)}
                  className={index === activeFactor ? S.active : ""}
                >
                  {factor}
                </li>
              );
            })}
          </ul>
          <input
            type="button"
            value="Применить"
            onClick={() => setActive(false)}
          />
        </div>
      </div>
    </div>
  );
};

export default FilterModal;
