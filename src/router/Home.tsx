import {useEffect, useState} from "react";
import {DarkModeToggle} from "../components/DarkModeToggle.tsx";
import "./Home.scss";
import {useNavigate} from "react-router-dom";

export function Home() {
    const [visible, setVisible] = useState(false);
    const [uuid, setUuid] = useState<string>(crypto.randomUUID());
    const navigate = useNavigate();

    useEffect(() => {
        const timer = setTimeout(() => setVisible(true), 100);
        return () => clearTimeout(timer);
    }, []);

    return (
        <>
            <DarkModeToggle/>
            <div className="content">
                <div className={`background ${visible ? "visible" : ""}`}>
                    <div className="title">
                        <img
                            className="logo"
                            src="/logo.png"
                            alt="QuickNote Logo"
                            draggable="false"
                        />
                    </div>

                    <div className="input-container">
                        <input
                            className="uuid-input"
                            type="text"
                            value={uuid}
                            onChange={(e) => setUuid(e.target.value)}
                        />
                        <button
                            className="submit-btn"
                            onClick={() => {
                                setVisible(false);
                                setTimeout(() => {
                                    navigate("/note/" + uuid);
                                }, 500);
                            }}
                        >
                            &rarr;
                        </button>
                    </div>
                </div>
            </div>
        </>
    );
}
