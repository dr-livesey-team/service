import React, {useEffect, useState} from "react";
import S from "./mainPage.module.scss";
import RequestBlock from "./RequestBlock/RequestBlock";
import SearchInput from "../../components/SearchInput/SearchInput";
import FilterModal from "../../components/FilterModal/FilterModal";
import CustomCheckbox from "../../components/CustomCheckbox/CustomCheckbox";
import MapModal from "../../components/MapModal/MapModal";

const MainPage = () => {
    const [filterActive, setFilterActive] = useState(false);
    const [mapActive, setMapActive] = useState(false);

    useEffect(() => {
        if (filterActive || mapActive) {
            document.body.style.overflow = "hidden";
            window.scrollTo(0,0);
        }
        else
            document.body.style.overflow = "";
    }, [filterActive, mapActive])
    return (
      <div className={S.mainPage}>
        <div className={S.header}>
          <div className={S.title}>
            <h2 className={S.title__header}>Заявки</h2>
            <p className={S.title__count}>3000 шт</p>
          </div>
          <div className={S.header__wrapper}>
            <SearchInput />
            <div className={S.sort} onClick={() => setFilterActive(true)}></div>
            <CustomCheckbox>Только аномальные</CustomCheckbox>
          </div>
        </div>
        <div className={S.content}>
            <div className={S.content__map}>
                <div className={S.mapButton} onClick={() => setMapActive(true)}>Найти на карте</div>
            </div>
          <RequestBlock />
        </div>
          {filterActive && <FilterModal setActive={setFilterActive}/>}
          {mapActive && <MapModal setActive={setMapActive} setFilterActive={setFilterActive}/>}
      </div>
    );
}

export default MainPage;
