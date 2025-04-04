import { useState, useEffect } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import { Chessboard } from "react-chessboard";
import { Chess, Move } from "chess.js";

const Game: React.FC = () => {
    const playerID: Number = Number(sessionStorage.getItem("playerID"));
    const gameID: Number = Number(sessionStorage.getItem("gameID"));

    const socketUrl = 'ws://localhost:8080/ws/game';
    const [messageHistory, setMessageHistory] = useState<MessageEvent<any>[]>([]);
    const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl, {
        shouldReconnect: () => true,
        reconnectAttempts: 10,
        reconnectInterval: 3000,
    });

    useEffect(() => {
        if (readyState === ReadyState.OPEN) {
            const joinMessage = {
              type: 'join',
              game_id: Number(gameID),
              player_id: Number(playerID),
              Content: "placeholder_name",
            };

            const json = JSON.stringify(joinMessage);
            console.log("Sending: ", json);
            sendMessage(json);
        }
    }, [readyState]);

    useEffect(() => {
        if (lastMessage == null) {
            return;
        }

        const messageData = JSON.parse(lastMessage.data);
        setMessageHistory((prev) => prev.concat(messageData));

        if (messageData.type === 'join') {
            if (messageData.player_id === playerID) {
                const color = messageData.content === "true" ? "white" : "black";
                setColor(color);
            }
        } else if (messageData.type === 'move') {
            const serverMove = messageData.content;
            const move = makeAMove(serverMove);
            if (!move) {
                console.log("Invalid move from server: ", serverMove);
            }
        } 
    }, [lastMessage]);

    const [color, setColor] = useState<"white" | "black">("white");
    const [game, setGame] = useState(new Chess());

    function makeAMove(move: { from: string; to: string; promotion?: string }): Move | null {
        const gameCopy = new Chess(game.fen());
        const result = gameCopy.move(move);

        if (result) {
            setGame(gameCopy);
        }

        return result;
    }

    function isDraggablePiece({ piece }: { piece: string }): boolean {
        return piece.charAt(0) === color.charAt(0);
    }

    function onDrop(sourceSquare: string, targetSquare: string): boolean {
        const move = makeAMove({
        from: sourceSquare,
        to: targetSquare,
        promotion: "q",
        });

        if (!move) {
            return false;
        }

        const messageObject = {
            type: 'move',
            game_id: gameID,
            player_id: playerID,
            content: convertToServerMove(move),
        };

        const json = JSON.stringify(messageObject)
        console.log("Sending: ", json);

        sendMessage(json);

        return true;
    }

    return (
        <div>
            <div>Game ID: {String(gameID)}</div>
            <Chessboard position={game.fen()} onPieceDrop={onDrop} isDraggablePiece={isDraggablePiece} boardOrientation={color} />
            {lastMessage ? <div>Last message: {lastMessage.data}</div> : null}
            <h2>debug message history</h2>
            <ul>
                {messageHistory.map((message, idx) => (
                    <li key={idx}>{message ? message.data : null}</li>
                ))}
            </ul>
        </div>
    );
};

function convertToServerMove(move: Move): string {
    let notation: string = move.piece.charAt(0)
    notation += move.before
    notation += move.after
    console.log(notation.toLowerCase())
    return notation.toLowerCase();
}

export default Game;