import React, { useState, useRef, useEffect, useCallback } from "react";

import "./app.css"

const AppWs = () => {
    const [isPaused, setIsPaused] = useState(false);
    const [data, setData] = useState(null);
    const [winMessage, setWinMesseg] = useState("");
    const [status, setStatus] = useState("");
    const ws = useRef(null);

    

    const request_to_gen = {
        method: "GenerateNumber"
    }

    const request_to_get = {
        method: "GetNumberByToken",
        body: {
            token: localStorage.getItem('token')
        }
    }

    //localStorage.removeItem('token');

    useEffect(() => {
        if (!isPaused) {
            ws.current = new WebSocket("ws://83.97.105.73:8080/ws"); // создаем ws соединение
            ws.current.onopen = () => {
                setStatus("Соединение открыто");  // callback на ивент открытия соединения
                console.log("connected");
                if(localStorage.getItem('token')){
                    ws.current.send(JSON.stringify(request_to_get));
                }else{
                    ws.current.send(JSON.stringify(request_to_gen));
                }
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
            if (message.status == "success"){
                if(message.message == "number"){
                    setData(message.data);
                }else if(message.message == "win"){
                    setWinMesseg("пойди туда - не знаю куда, \t возьми то - не знаю что пвыпвапвапвп вапвы пвап вап вап вап выпвапв пвп впвп ");
                }
            }else{
                if(message.message == "completed"){
                    setData({completed:"Лотерея завершена"})
                }
                else if(message.message == "token is invalid"){
                    ws.current.send(JSON.stringify(request_to_gen));
                }
            }
            
        };
    }, [isPaused]);
    if(data && data.token){
        localStorage.setItem('token', data.token);
    }

    return (
        <>
            <div class="text-block">
            <img className="image" src="img/PTU.png"/>
            {data && data.number &&<div className="box">
                        <p className="text1">{`${data && data.number}`}</p>
                </div>}
                {data && data.completed &&<div className="box2">
                        <p className="text2">{`${data && data.completed}`}</p>
                </div>}
                {winMessage != "" &&
                    
                    <div className="box2">
                            <p className="text2">{`${winMessage}`}</p>
                    </div>
                }
            </div>
            {!!data &&
                <div>
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
            }
        </>
    )
}

/*
    <button onClick={() => {
                        ws.current.close();
                        setIsPaused(!isPaused)
                    }}>{!isPaused ? 'Остановить соединение' : 'Открыть соединение' }</button>
*/

export default AppWs;