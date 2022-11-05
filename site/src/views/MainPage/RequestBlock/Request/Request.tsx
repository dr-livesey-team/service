import React from "react";
import S from "./request.module.scss";
import classnames from "classnames";

const Request:React.FC<{minRequest?: boolean}> = ({minRequest}) => {
    return <tr className={classnames(S.request, minRequest ? S.minimal : '')}>
      <td className={S.abnormal}>Да</td>
      <td>Разбито/сломано/повреждено окно  в местах общего пользования</td>
      <td>31.08.2022, 18:22</td>
      {!minRequest && <td>ГБУ “Жилищник района Теплый стан”</td>}
      {!minRequest && <td>ПАО "МОЭК" </td>}
    </tr>;
}

export default Request;
