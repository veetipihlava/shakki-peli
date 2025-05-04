import React, { useState, useEffect } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import { Chessboard } from "react-chessboard";
import { Chess, Move } from "chess.js";
import Navbar from './Navbar.tsx';
import { useNavigate } from 'react-router-dom';
import './Game.css';

const Game: React.FC = () => {
    const userID: number = Number(sessionStorage.getItem("userID"));
    const gameID: number = Number(sessionStorage.getItem("gameID"));
    const navigate = useNavigate();

    const socketUrl = 'ws://localhost:8080/ws';
    const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl, {
        shouldReconnect: () => true,
        onError: (event) => {
            console.error("WebSocket error: ", event);
        },
        onOpen: () => {
            console.log("WebSocket connection opened");
        }
    });

    const [color, setColor] = useState<"white" | "black">("white");
    const [game, setGame] = useState(new Chess());
    const [previousState, setPreviousState] = useState(new Chess());
    const [kingInCheck, setKingInCheck] = useState(false);
    const [timer, setTimer] = useState(0);
    const [opponentJoined, setOpponentJoined] = useState(false);
    const [winner, setWinner] = useState<string | null>(null);
    const [isDraw, setIsDraw] = useState(false);

    useEffect(() => {
        if (readyState === ReadyState.OPEN) {
            const joinMessage = {
                type: 'join',
                game_id: gameID,
                user_id: userID,
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

        const messageData = JSON.parse(lastMessage.data);
        console.log("Received: ", messageData);

        if (messageData.type === "join") {
            if (messageData.content.user_id !== userID) {
                setOpponentJoined(true);
            } else if (messageData.content.user_id === userID) {
                setOpponentJoined(true);
                const color = messageData.content.color === true ? "white" : "black";
                setColor(color);
            }
        } else if (messageData.type === 'move') {
            const serverMove = convertFromServerMove(messageData.content.move);

            if (messageData.content.king_in_check) {
                setKingInCheck(true);
            } else {
                setKingInCheck(false);
            }

            if (messageData.content.checkmate) {
                const winnerColor = messageData.content.winner_color ? "White" : "Black";
                setWinner(winnerColor); // Set the winner
            }

            if (messageData.content.draw) {
                setIsDraw(true); // Set the draw state
            }

            if (messageData.content.user_id !== userID) {
                const move = makeAMove(serverMove);
                if (!move) {
                    console.log("Invalid move from server: ", serverMove);
                }
            }
        } else if (messageData.type === 'closing') {
            alert("Game closed by opponent");
            navigate('/');
        } else if (messageData.type === 'error' && messageData.content.ref_type === "move") {
            setGame(previousState);
        }
    }, [lastMessage]);

    useEffect(() => {
        let interval: NodeJS.Timeout | null = null;

        if (opponentJoined) {
            interval = setInterval(() => {
                setTimer((prev) => prev + 1);
            }, 1000);
        }

        return () => {
            if (interval) clearInterval(interval);
        };
    }, [opponentJoined]);

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

        const json = JSON.stringify(messageObject);
        console.log("Sending: ", json);

        sendMessage(json);

        return true;
    }

    function handleLeaveGame() {
        const leaveMessage = {
            type: 'leave',
            game_id: gameID,
            user_id: userID,
        };

        const json = JSON.stringify(leaveMessage);
        console.log("Sending: ", json);
        sendMessage(json);

        // Navigate back to the home or join page
        navigate('/');
    }

    const formatTime = (timeInSeconds: number): string => {
        const minutes = Math.floor(timeInSeconds / 60);
        const seconds = timeInSeconds % 60;
        return `${minutes}:${seconds < 10 ? `0${seconds}` : seconds}`;
    };

    return (
        <div>
            <Navbar />
            <div className='game-window-container'>
                <div className='game-info-container'>
                    <div className='game-invite-container'>
                        <div>Game ID: {String(gameID)}</div>
                    </div>
                    <div className='game-invite-container'>
                        {opponentJoined && <div className='timer'>⏱️ Time Elapsed: {formatTime(timer)}</div>}
                    </div>
                </div>
                <div className='game-board-container'>
                    <Chessboard position={game.fen()} onPieceDrop={onDrop} isDraggablePiece={isDraggablePiece} boardOrientation={color} boardWidth={500} />
                </div>
                <div className='game-message-container'>
                    {kingInCheck && <div >
                        <div className='check-warning'>⚠️ Your king is in check!</div>
                    </div>}
                </div>
                <button className='leave-game-button' onClick={handleLeaveGame}>
                    Leave Game
                </button>
            </div>
            {winner && (
                <div className='overlay'>
                    <div className='overlay-card'>
                        <h2>Game Over</h2>
                        <p>{winner} wins!</p>
                        <button onClick={() => navigate('/')}>Return to Home</button>
                    </div>
                </div>
            )}
            {isDraw && (
                <div className='overlay'>
                    <div className='overlay-card'>
                        <h2>Game Over</h2>
                        <p>The game is a draw!</p>
                        <button onClick={() => navigate('/')}>Return to Home</button>
                    </div>
                </div>
            )}
        </div>
    );
};

function convertToServerMove(move: Move): string {
    let notation: string = move.piece.charAt(0);
    notation += move.from;
    notation += move.to;
    return notation.toUpperCase();
}

function convertFromServerMove(serverMove: string): { from: string; to: string; promotion?: string } {
    const lowerCaseMove = serverMove.toLowerCase();
    const from = lowerCaseMove.slice(1, 3);
    const to = lowerCaseMove.slice(3, 5);

    return { from, to, promotion: "q" };
}

export default Game;