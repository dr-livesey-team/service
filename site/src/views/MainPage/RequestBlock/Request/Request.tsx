import React from "react";
import S from "./request.module.scss";
import classnames from "classnames";
import {Link} from "react-router-dom";

const Request: React.FC<{ minRequest?: boolean; anomaly: any; isActive?: boolean }> = ({ minRequest, anomaly, isActive }) => {
  return (
    <Link to={'/request?id=' + anomaly?.id} id={'anom' + anomaly?.id} className={classnames(S.request, minRequest ? S.minimal : "", isActive ? S.active : "")}>
      <td>{anomaly?.number}</td>
      <td>{anomaly?.fault_name}</td>
      <td>{anomaly?.opening_date}</td>
      {!minRequest && <td>{anomaly?.management_company_name}</td>}
      {!minRequest && <td>{anomaly?.service_organization_name}</td>}
    </Link>
  );
};

export default Request;
