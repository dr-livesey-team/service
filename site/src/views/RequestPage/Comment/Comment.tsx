import React from "react";
import S from "./comment.module.scss";

const Comment: React.FC<{item: any}> = ({item}) => {
  return (
    <tr>
      <td>{item?.request_root_identifier}</td>
      <td className={S.date}>{item?.opening_date?.slice(0, 10)}</td>
      <td className={S.date}>{item?.closing_date?.slice(0, 10)}</td>
      <td>{item?.effectiveness}</td>
      <td>{item?.fault_name}</td>
      <td className={`${S.review}`}>Смотреть отзывы</td>
    </tr>
  );
};

export default Comment;
