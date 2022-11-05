import React, {Component} from 'react';
import Request from "./Request/Request";
import S from './requestBlock.module.scss'

class RequestBlock extends Component {
    render() {
        return (
            <table className={S.requestBlock}>
                <th className={S.header}>
                    <td>Аномальная</td>
                    <td>Наименование</td>
                    <td>Дата создания</td>
                    <td>Управляющая копания</td>
                    <td>Обслуживающая организация</td>
                </th>
                <Request/>
                <Request/>
                <Request/>
                <Request/>
            </table>
        );
    }
}

export default RequestBlock;
