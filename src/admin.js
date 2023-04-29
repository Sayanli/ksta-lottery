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
            ws.current = new WebSocket("ws://192.168.1.84:8080/ws"); // создаем ws соединение
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
            <div class="text-block">
                    <button className="button" onClick={() => {
                            console.log("completed");
                            ws.current.send(JSON.stringify(request_to_completed));
                        }}>
                        <p className="text2">Завершить</p>
                    </button>
            </div>
            <div>
                    
                    <div class="area" >
                    <ul class="circles">
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
                            <li></li>
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