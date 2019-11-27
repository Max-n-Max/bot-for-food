'use strict';

app.service("SocketService",["apiClient",
    function(apiClient){

        /*
            WebSocket Event Connecting
        */
        var CONNECTING = WebSocket.CONNECTING;

        /*
            WebSocket Event Open
        */
        var OPEN  = WebSocket.OPEN;

        /*
            WebSocket Event Closing
        */
        var CLOSING  = WebSocket.CLOSING;

        /*
            WebSocket Event Closed
        */
        var CLOSED  = WebSocket.CLOSED;

        /*
            Timeout in seconds
        */
        var TIMEOUT_STAMP = 3;

        /*
            Maximum attempts from last success connection (or first initiliaze)
        */
        var MAX_ATTEMPS = 3;

        /*
            WebSocket protocol that server listens to
        */
        var SOCKET_PROTOCOL = "echo-protocol";

        /*
            WebSocket Instance
        */
        var connection  = null;

        /*
           Current attemp
        */
        var currentAttemp = 0;

        /*
            didOpen, indicates if open \ closed
        */
        var didOpen  = false;

        /*
             is debug mode
        */
        var isDebug = true;
        /*=================METHODS===============*/
        (function(w){
            //firefox
            if(w.MozWebSocket) w.WebSocket = MozWebSocket;
        }(window));

        var supportsWebSocket = function(){
            return "WebSocket" in window;
        };

        var close = function(){
            didOpen = false;
            if(connection){
                wsLogger(console.log,"WebSocket :: close connection");
                connection.close();
            }
        };

        var open = function(){

            console.log("WebSocket :: open is called");

            currentAttemp = 0;
            didOpen = true;
            init();
        };

        var isHttps = function(){
            return /https/.test(document.location.protocol);
        }

        var generateSocketPath = function(){
            return [
                "ws://", //Protocol
                document.location.hostname, //Host
                ":9092",   //Port
                "/ws", //method
                "?lastMod="  //URL query for hashKey
            ].join("");
        };

        var setConnection = function(hash){

            if(connection) return;

            console.log("will set Connection WebSocket::" + generateSocketPath() + hash);

            connection = new WebSocket(generateSocketPath() + hash/*,SOCKET_PROTOCOL*/);

            connection.onerror = function(err){
                console.error(err);
                //setTimeout(init,TIMEOUT_STAMP * 1000);
            };

            connection.onopen = function(event){
                console.log("WebSocket connection Succesful!");
                currentAttemp = 0;
            };

            connection.onmessage = function(event){
                try{
                    console.log("DID RECEIVE MSG FROM SEB SOCKET:\n==================\n" + event.data + "\n\n==END WEB SOCKET MSG===");
                    var json = JSON.parse(event.data);

                }catch(err){
                    console.error(err);
                }
            };

            connection.onclose = function(event){
                console.log("WebSocket connection closed!");
                connection = null;
                //setTimeout(init,TIMEOUT_STAMP * 1000);
            };
        };

        var init = function(){

            if(!didOpen){
                return;
            };

            if(++currentAttemp > MAX_ATTEMPS){
                return;
            }

            if(!supportsWebSocket()){
                return;
            }

            console.log("WebSocket :: socketRequestTicket ...");

            setConnection("15d39d2df32d7e35");

            // apiClient.socketRequestToken().then(function(res){
            //     console.log("WebSocket :: socketRequestTicket - responose!" + JSON.stringify(res));
            //         if(res.result && res.result === "success" && res.data){
            //
            //             wsLogger(console.log,"WebSocket :: socketRequestTicket - success!");
            //
            //             var subResult = res.data;
            //
            //             if(subResult.result && subResult.result === "success" && subResult.data && subResult.data.ticket_hash)
            //                 setConnection(subResult.data.ticket_hash);
            //         }
            //     })
            //     .catch(function(err){
            //         setTimeout(init,3000);
            //     });
        };


        return {
            open  : open,
            close : close
        };

    }]);
