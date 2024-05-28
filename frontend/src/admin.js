import React, { useState, useRef, useEffect, useCallback } from "react";

import "./app.css"

const Admin = () => {
    const [isPaused, setIsPaused] = useState(false);
    const [data, setData] = useState(null);
    const [status, setStatus] = useState("");
    const ws = useRef(null);

    

    const request_to_completed = {
        method: "LotteryCompleted"
    }

    useEffect(() => {
        if (!isPaused) {
            ws.current = new WebSocket("ws://83.97.105.73:8080/ws"); // создаем ws соединение
            ws.current.onopen = () => {
                setStatus("Соединение открыто");  // callback на ивент открытия соединения
                console.log("connected");
                console.log("compete");
            }
            ws.current.onclose = () => setStatus("Соединение закрыто"); // callback на ивент закрытия соединения

            gettingData();

            
            
        }

        return () => ws.current.close(); // кода меняется isPaused - соединение закрывается
    }, [ws, isPaused]);

    const gettingData = useCallback(() => {
        if (!ws.current) return;
        console.log("getting");
        ws.current.onmessage = e => {                //подписка на получение данных по вебсокету
            if (isPaused) return;
            const message = JSON.parse(e.data);
            console.log(message);
        };
    }, [isPaused]);

    return (
        <>
            <div>
            <button className="button" onClick={() => {
                            console.log("completed");
                            ws.current.send(JSON.stringify(request_to_completed));
                        }}>
                        Завершить
            </button>
                    <div class="area" >
                    <ul class="circles">
                    <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Круглешок.png"/></li>
                            <li><img className="img-figure" src="img/Мозги.png"/></li>
                            <li><img className="img-figure" src="img/Мозги.png"/></li>
                            <li><img className="img-figure" src="img/Мозги.png"/></li>
                    </ul>
                    </div >
            </div>
            
        </>
    )
}

/*
    <button onClick={() => {
                        ws.current.close();
                        setIsPaused(!isPaused)
                    }}>{!isPaused ? 'Остановить соединение' : 'Открыть соединение' }</button>
*/

export default Admin;