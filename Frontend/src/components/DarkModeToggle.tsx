import {useRef, useState} from "react";
import {useDarkMode} from "../hooks/useDarkMode";
import "./DarkMode.scss";

type Ripple = {
  key: number;
  x: number;
  y: number;
};

export const DarkModeToggle = () => {
  const { isDarkMode, toggleDarkMode } = useDarkMode();
  const [ripples, setRipples] = useState<Ripple[]>([]);
  const buttonRef = useRef<HTMLButtonElement>(null);

  const handleClick = () => {
    if (buttonRef.current) {
      const rect = buttonRef.current.getBoundingClientRect();
      const centerX = rect.left + rect.width / 2;
      const centerY = rect.top + rect.height / 2;

      setRipples((prevRipples) => [
        ...prevRipples,
        {
          key: Date.now(),
          x: centerX,
          y: centerY
        }
      ]);
    }
    toggleDarkMode();
  };

  const handleAnimationEnd = (key: number) => {
    setRipples((prevRipples) => prevRipples.filter((r) => r.key !== key));
  };

  return (
    <>
      <button
        ref={buttonRef}
        className="theme-toggle"
        onClick={handleClick}
        aria-label={isDarkMode ? "Activate light mode" : "Activate dark mode"}
      >
        <div className="icon-container">
          <svg
            className="sun-and-moon"
            aria-hidden="true"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <mask className="moon-mask" id="moon-mask">
              <rect x="0" y="0" width="100%" height="100%" fill="white" />
              <circle cx="24" cy="10" r="5" fill="black" />
            </mask>
            <circle
              className="sun"
              cx="12"
              cy="12"
              r="6"
              mask="url(#moon-mask)"
              fill="currentColor"
            />
            <g className="sun-beams" stroke="currentColor">
              <line x1="12" y1="1" x2="12" y2="3" />
              <line x1="12" y1="21" x2="12" y2="23" />
              <line x1="4.22" y1="4.22" x2="5.64" y2="5.64" />
              <line x1="18.36" y1="18.36" x2="19.78" y2="19.78" />
              <line x1="1" y1="12" x2="3" y2="12" />
              <line x1="21" y1="12" x2="23" y2="12" />
              <line x1="4.22" y1="19.78" x2="5.64" y2="18.36" />
              <line x1="18.36" y1="5.64" x2="19.78" y2="4.22" />
            </g>
          </svg>
        </div>
      </button>
      {ripples.map(({ key, x, y }) => (
        <span
          key={key}
          className="ripple"
          style={{ left: `${x}px`, top: `${y}px` }}
          onAnimationEnd={() => handleAnimationEnd(key)}
        />
      ))}
    </>
  );
};
