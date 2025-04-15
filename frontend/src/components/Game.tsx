import { useState, useEffect } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import { Chessboard } from "react-chessboard";
import { Chess, Move } from "chess.js";

const Game: React.FC = () => {
    const userID: Number = Number(sessionStorage.getItem("userID"));
    const gameID: Number = Number(sessionStorage.getItem("gameID"));

    const socketUrl = 'ws://localhost:8080/ws';
    const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl, {
        shouldReconnect: () => true,
        // reconnectAttempts: 10,
        // reconnectInterval: 3000,
    });

    useEffect(() => {
        if (readyState === ReadyState.OPEN) {
            const joinMessage = {
              type: 'join',
              game_id: Number(gameID),
              user_id: Number(userID),
              Content: "placeholder_name",
            };

            const json = JSON.stringify(joinMessage);
            console.log("Sending: ", json);
            sendMessage(json);
        } else {
            console.log(readyState)
        }
    }, [readyState]);

    useEffect(() => {
        if (!lastMessage) return;

        const messageData = JSON.parse(lastMessage.data)
        console.log("Recieved: ", messageData);

        if (messageData.type === "join") {
            if (messageData.content.user_id === userID) {
                const color = messageData.content.color === true ? "white" : "black";
                setColor(color);
            }
        } else if (messageData.type === 'move') {
            const serverMove = convertFromServerMove(messageData.content.move);
            if (messageData.content.user_id !== userID) {
                const move = makeAMove(serverMove);
                if (!move) {
                    console.log("Invalid move from server: ", serverMove);
                }
            }
        } else if (messageData.type === 'error' && messageData.content.ref_type == "move") {
            setGame(previousState);
            console.log("reverting move")
        }
    }, [lastMessage]);

    const [color, setColor] = useState<"white" | "black">("white");
    const [game, setGame] = useState(new Chess());
    const [previousState, setPreviousState] = useState(new Chess());

    function makeAMove(move: { from: string; to: string; promotion?: string }): Move | null {
        const gameCopy = new Chess(game.fen());
        setPreviousState(new Chess(gameCopy.fen()));
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
            user_id: userID,
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
            <Chessboard position={game.fen()} onPieceDrop={onDrop} isDraggablePiece={isDraggablePiece} boardOrientation={color} boardWidth={500} />
            {lastMessage ? <div>Last message: {lastMessage.data}</div> : null}
        </div>
    );
};

function convertToServerMove(move: Move): string {
    let notation: string = move.piece.charAt(0)
    notation += move.from
    notation += move.to
    return notation.toUpperCase();
}

function convertFromServerMove(serverMove: string): { from: string; to: string; promotion?: string } {
    const lowerCaseMove = serverMove.toLowerCase();
    const from = lowerCaseMove.slice(1, 3)
    const to = lowerCaseMove.slice(3, 5)

    return { from, to, promotion: "q" };
}

export default Game;