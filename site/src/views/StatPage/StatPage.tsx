import React, { useEffect, useState } from "react";
import {
  AreaChart,
  Area,
  Brush,
  CartesianGrid,
  Line,
  LineChart,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";
import S from "./statPage.module.scss";
import axios from "axios";
import { pointsMock, serverURL } from "../../assets/requestMock";
import {
  anomaly_category_list,
  district_name_list,
  management_company_name_list,
} from "../../assets/filterLists";

const StatPage: React.FC = () => {
  const [graphArray, setGraphArray] = useState([1]);
  const [anomalyCategory, setAnomalyCategory] = useState<any>("");
  const [districtName, setDistrictName] = useState<any>("");
  const [managementCompanyName, setManagementCompanyName] = useState<any>("");
  const [period, setPeriod] = useState<any>("");
  const [points, setPoints] = useState<any>([]);
  const addGraph = () => {
    setGraphArray([...graphArray, graphArray[graphArray.length - 1]]);
  };

  const refactorDate = (date: string) => {
    let today = new Date(Date.now());
    switch (date) {
      case "За последний день":
        return (
          today.getFullYear() +
          "-" +
          (today.getMonth() < 10 ? "0" + today.getMonth() : today.getMonth()) +
          "-" +
          (today.getDate() - 1 < 10
            ? "0" + (today.getDate() - 1)
            : today.getDate() - 1)
        );
      case "За последний месяц":
        return (
          today.getFullYear() +
          "-" +
          (today.getMonth() - 1 < 10
            ? "0" + (today.getMonth() - 1)
            : today.getMonth() - 1) +
          "-" +
          (today.getDate() < 10 ? "0" + today.getDate() : today.getDate())
        );
      case "За последний год":
        return (
          today.getFullYear() -
          1 +
          "-" +
          (today.getMonth() < 10 ? "0" + today.getMonth() : today.getMonth()) +
          "-" +
          (today.getDate() < 10 ? "0" + today.getDate() : today.getDate())
        );
      default:
        return "";
    }
  };

  useEffect(() => {
    let today = new Date(Date.now());
    axios
      .get(
        serverURL +
          "statistic?" +
          "opening_date=" +
          refactorDate(period) +
          "&closing_date=" +
          today.getFullYear() +
          "-" +
          (today.getMonth() < 10 ? "0" + today.getMonth() : today.getMonth()) +
          "-" +
          (today.getDate() < 10 ? "0" + today.getDate() : today.getDate()) +
          "&district_name=" +
          districtName +
          "&management_company_name=" +
          managementCompanyName +
          "&anomaly_category_name=" +
          anomalyCategory
      )
      .then((response) => setPoints(response.data?.points))
      .catch(() => setPoints(pointsMock));
  }, [anomalyCategory, districtName, managementCompanyName, period]);
  return (
    <div className={S.page__wrapper}>
      <div className={S.header}>
        <h2 className={S.title}>Статистика</h2>
      </div>
      <div className={S.content}>
        {graphArray.map((val, index) => (
          <>
            <h2>Доля ситуаций с аномалией</h2>
            <div className={S.graph}>
              <div className={S.graph__left}>
                <LineChart
                  width={window.innerWidth > 1200 ? 700 : 600}
                  height={460}
                  data={points}
                  margin={{ top: 40, right: 40, bottom: 20, left: 20 }}
                >
                  <CartesianGrid vertical={false} />
                  <XAxis dataKey="date" />
                  <YAxis domain={["auto", "auto"]} />
                  <Tooltip
                    wrapperStyle={{
                      borderColor: "white",
                      boxShadow: "2px 2px 3px 0px rgb(204, 204, 204)",
                    }}
                    contentStyle={{
                      backgroundColor: "rgba(255, 255, 255, 0.8)",
                    }}
                    labelStyle={{ fontWeight: "bold", color: "#666666" }}
                  />
                  <Line
                    dataKey="percent"
                    type="monotone"
                    stroke="#ff7300"
                    dot={false}
                  />
                  <Brush dataKey="date" startIndex={points.length - 40}>
                    <AreaChart>
                      <CartesianGrid />
                      <YAxis hide domain={["auto", "auto"]} />
                      <Area
                        dataKey="percent"
                        stroke="#ff7300"
                        fill="#ff7300"
                        dot={false}
                      />
                    </AreaChart>
                  </Brush>
                </LineChart>
              </div>
              <div className={S.graph__right}>
                <ul>
                  <li>
                    <label>Вид аномалии</label>
                    <select
                      value={anomalyCategory}
                      onChange={(e) => setAnomalyCategory(e.target.value)}
                    >
                      {anomaly_category_list.map((item) => (
                        <option>{item}</option>
                      ))}
                    </select>
                  </li>
                  <li>
                    <label>Период</label>
                    <select
                      value={period}
                      onChange={(e) => setPeriod(e.target.value)}
                    >
                      <option>За весь период</option>
                      <option>За последний день</option>
                      <option>За последний месяц</option>
                      <option>За последний год</option>
                    </select>
                  </li>
                  <li>
                    <label>Район</label>
                    <select
                      value={districtName}
                      onChange={(e) => setDistrictName(e.target.value)}
                    >
                      {district_name_list.map((item) => (
                        <option>{item}</option>
                      ))}
                    </select>
                  </li>
                  <li>
                    <label>Управляющая компания</label>
                    <select
                      value={managementCompanyName}
                      onChange={(e) => setManagementCompanyName(e.target.value)}
                    >
                      {management_company_name_list.map((item) => (
                        <option>{item}</option>
                      ))}
                    </select>
                  </li>
                </ul>
              </div>
            </div>
            <div className={S.button__wrapper}>
              {index === graphArray.length - 1 ? (
                <input value="Добавить график" onClick={() => addGraph()} />
              ) : (
                ""
              )}
              <input className={S.button_accent} value="Распечатать график" />
            </div>
          </>
        ))}
      </div>
    </div>
  );
};

export default StatPage;
