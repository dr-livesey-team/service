import React, { Component } from "react";
import S from "./header.module.scss";
import { NavLink } from "react-router-dom";

class Header extends Component {
  render() {
    return (
      <div className={S.header}>
        <div className={S.logo}>Ливси Х</div>
        <div className={S.menu}>
          <ul>
            <NavLink
              to="/statistics"
              title="Подраздел"
              className={(isActive) => (isActive ? S.active : "")}
            >
              Статистика
            </NavLink>
            <NavLink
              to="/monitoring"
              title="Подраздел"
              className={(isActive) => (isActive ? S.active : "")}
            >
              Мониторинг
            </NavLink>
          </ul>
        </div>
        <div className={S.profile}>
          <svg
            width="16"
            height="20"
            viewBox="0 0 16 20"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            className={S.icon}
          >
            <path
              fillRule="evenodd"
              clipRule="evenodd"
              d="M6 2C6 0.896 6.896 0 8 0C9.104 0 10 0.896 10 2V2.3373C12.4831 3.207 14.2222 5.65252 14.2222 8.382V10.9055C14.2222 11.3657 14.408 11.8076 14.7396 12.134L15.5822 12.9635C15.8498 13.2268 16 13.5847 16 13.9566V15.221C16 15.6515 15.6462 15.9997 15.2089 15.9997H0.791111C0.353778 15.9997 0 15.6515 0 15.221V13.9566C0 13.5847 0.150222 13.2268 0.417778 12.9635L1.26044 12.134C1.592 11.8076 1.77778 11.3657 1.77778 10.9055V8.12475C1.77778 5.43104 3.54474 3.1434 6 2.32364V2ZM9 2.01068H8.36162C8.30501 2.00743 8.24817 2.00497 8.19111 2.0033C8.00461 1.99778 7.81989 2.00035 7.63728 2.01068H7V2C7 1.44828 7.44828 1 8 1C8.55172 1 9 1.44828 9 2V2.01068ZM5.01592 16.529C4.75577 16.6216 4.61995 16.9075 4.71255 17.1677C5.20651 18.5553 6.50987 19.5 7.9996 19.5C9.49016 19.5 10.7936 18.5554 11.2876 17.1677C11.3803 16.9075 11.2444 16.6216 10.9843 16.529C10.7241 16.4363 10.4382 16.5722 10.3456 16.8323C9.99173 17.8263 9.062 18.5 7.9996 18.5C6.93808 18.5 6.00843 17.8262 5.65464 16.8323C5.56204 16.5722 5.27607 16.4363 5.01592 16.529ZM15 13.9566V14.9997H1V13.9566C1 13.8529 1.04255 13.7517 1.11929 13.6761L1.96196 12.8466C2.48382 12.3329 2.77778 11.6346 2.77778 10.9055V8.12475C2.77778 5.23932 5.20623 2.91553 8.16157 3.00286C10.9675 3.08505 13.2222 5.50825 13.2222 8.382V10.9055C13.2222 11.6346 13.5162 12.3329 14.038 12.8466L14.8807 13.6761C14.9575 13.7517 15 13.8529 15 13.9566Z"
              fill="black"
            />
          </svg>
          Иванова М.П.
        </div>
      </div>
    );
  }
}

export default Header;
