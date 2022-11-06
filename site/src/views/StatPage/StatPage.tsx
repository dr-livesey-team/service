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
import {pointsMock, serverURL} from "../../assets/requestMock";

const StatPage: React.FC = () => {
  const [graphArray, setGraphArray] = useState([1]);
  const [points, setPoints] = useState<any>([]);
  const addGraph = () => {
    setGraphArray([...graphArray, graphArray[graphArray.length - 1]]);
  };

  useEffect(() => {
    axios
      .get(serverURL + "statistic")
      .then((response) => setPoints(response.data?.points))
      .catch(() => setPoints(pointsMock));
  }, []);
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
                    <select>
                      <option>Все</option>
                      <option>1</option>
                      <option>2</option>
                    </select>
                  </li>
                  <li>
                    <label>Период</label>
                    <select>
                      <option>За весь период</option>
                      <option>1</option>
                      <option>2</option>
                    </select>
                  </li>
                  <li>
                    <label>Район</label>
                    <select>
                      <option>Все</option>
                      <option>1</option>
                      <option>2</option>
                    </select>
                  </li>
                  <li>
                    <label>Управляющая компания</label>
                    <select>
                      <option>Все</option>
                      <option>1</option>
                      <option>2</option>
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
