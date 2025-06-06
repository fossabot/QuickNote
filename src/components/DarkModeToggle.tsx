import {useDarkMode} from '../hooks/useDarkMode';
import './DarkMode.scss';
import {useRef} from 'react';

export function DarkModeToggle() {
    const {isDarkMode, toggleDarkMode} = useDarkMode();
    const buttonRef = useRef<HTMLButtonElement>(null);

    const handleClick = () => {
        const button = buttonRef.current;
        if (!button) return;

        const circle = document.createElement('span');
        circle.className = 'ripple';

        const rect = button.getBoundingClientRect();
        const x = rect.left + rect.width / 2;
        const y = rect.top + rect.height / 2;

        console.log(`ripple: x: ${x}, y: ${y}`);

        circle.style.left = `${x}px`;
        circle.style.top = `${y}px`;

        document.body.appendChild(circle);

        circle.addEventListener('animationend', () => {
            circle.remove();
        });

        toggleDarkMode();
    };

    return (
        <button
            className="theme-toggle"
            ref={buttonRef}
            onClick={handleClick}
        >
            {isDarkMode ? '☀️' : '🌙'}
        </button>
    );
}
