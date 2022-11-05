import React from "react";
import S from "./comment.module.scss";

const Comment = () => {
  return (
    <tr>
      <td className={S.date}>15:11, 20.10.2021</td>
      <td>10ч 10мин</td>
      <td>Выполнено</td>
      <td>Перезапуск системы</td>
      <td>В 1 комнате еле теплые батареи, а во 2 комнатее вообще нет тепла</td>
      <td className={`${S.review} ${S.active}`}>Смотреть отзывы</td>
    </tr>
  );
};

export default Comment;
