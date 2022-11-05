import React, { useState } from "react";
import S from "./filterItem.module.scss";

const FilterItem: React.FC<{ children: string }> = ({ children }) => {
  const [active, setActive] = useState(false);
  return (
    <li onClick={() => setActive(!active)} className={active ? S.active : ''}>
      <input checked={active} type="checkbox" className={S.customCheckbox} />
      <label className={S.label}>{children}</label>
    </li>
  );
};

export default FilterItem;
