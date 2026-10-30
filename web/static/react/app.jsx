const { useState, useEffect, useRef } = React;

const ALGORITHMS = [
    { value: "selection", label: "Selection Sort" },
    { value: "bubble", label: "Bubble Sort" },
    { value: "insertion", label: "Insertion Sort" },
    { value: "merge", label: "Merge Sort" },
    { value: "quick", label: "Quick Sort" },
    { value: "heap", label: "Heap Sort" },
    { value: "shell", label: "Shell Sort" },
    { value: "cocktail", label: "Cocktail Shaker Sort" },
];

const NAV_LINKS = [
    { key: "visualizer", href: "/", label: "Visualizer" },
    { key: "algorithms", href: "/algorithms", label: "Algorithms" },
    { key: "about", href: "/about", label: "About" },
    { key: "angular", href: "/angular", label: "Angular" },
];

function Nav() {
    const page = document.body.dataset.page;

    return (
        <nav className="site-nav">
            <a className="brand" href="/">
                🔮 SortViz
            </a>
            <div className="nav-links">
                {NAV_LINKS.map((link) => (
                    <a
                        key={link.key}
                        href={link.href}
                        className={page === link.key ? "active" : ""}
                    >
                        {link.label}
                    </a>
                ))}
            </div>
        </nav>
    );
}

function Modal({ title, onClose, children }) {
    return (
        <div className="popup" onClick={onClose}>
            <div className="popup-content" onClick={(event) => event.stopPropagation()}>
                <button className="close-btn" onClick={onClose} aria-label="Close">
                    ×
                </button>
                <h3>{title}</h3>
                {children}
            </div>
        </div>
    );
}

function Visualizer() {
    const [bars, setBars] = useState(() => Array.from({ length: 40 }, (_, i) => i + 1));
    const [algorithm, setAlgorithm] = useState("selection");
    const [speedValue, setSpeedValue] = useState(6);
    const [steps, setSteps] = useState(null);
    const [stepIndex, setStepIndex] = useState(-1);
    const [playing, setPlaying] = useState(false);
    const [status, setStatus] = useState("");
    const [modal, setModal] = useState(null);
    const timerRef = useRef(null);

    const delay = (11 - speedValue) * 100;

    useEffect(() => {
        if (!playing || !steps) {
            return undefined;
        }
        if (stepIndex >= steps.length - 1) {
            setPlaying(false);
            setStatus("Array sorted!");
            return undefined;
        }
        timerRef.current = setTimeout(() => {
            setStepIndex((index) => index + 1);
        }, delay);
        return () => clearTimeout(timerRef.current);
    }, [playing, steps, stepIndex, delay]);

    const displayBars = steps && stepIndex >= 0 ? steps[stepIndex].array : bars;

    const resetRun = () => {
        setPlaying(false);
        setSteps(null);
        setStepIndex(-1);
    };

    const shuffle = () => {
        const next = bars.slice();
        for (let i = next.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            [next[i], next[j]] = [next[j], next[i]];
        }
        setBars(next);
        resetRun();
        setStatus("");
    };

    const start = async () => {
        if (steps && stepIndex >= steps.length - 1) {
            setStatus("Array already sorted! Please shuffle and try again.");
            return;
        }
        if (steps) {
            setPlaying(true);
            return;
        }

        setStatus("");
        try {
            const response = await fetch("/sort", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ array: bars, algorithm }),
            });
            if (!response.ok) {
                throw new Error(response.statusText || "Request failed");
            }
            const data = await response.json();
            if (!Array.isArray(data) || data.length === 0) {
                throw new Error("No sort steps returned");
            }
            setSteps(data);
            setStepIndex(0);
            setPlaying(true);
        } catch (error) {
            setStatus("Error: " + error.message);
        }
    };

    const fetchAlgorithmData = async () => {
        const response = await fetch("/static/data/algorithm_data.json");
        if (!response.ok) {
            throw new Error("Could not load algorithm data");
        }
        const data = await response.json();
        return data[algorithm];
    };

    const showComplexity = async () => {
        try {
            const data = await fetchAlgorithmData();
            setModal({ title: "Complexity", text: data.complexity });
        } catch (error) {
            setStatus("Error: " + error.message);
        }
    };

    const showCode = async () => {
        try {
            const data = await fetchAlgorithmData();
            setModal({
                title: "Code",
                code: data.code.replace(/\{\{NEWLINE\}\}/g, "\n"),
            });
        } catch (error) {
            setStatus("Error: " + error.message);
        }
    };

    return (
        <div className="container">
            <aside className="sidebar">
                <h2>⚙️ Controls</h2>

                <div className="algorithm-select">
                    <select
                        value={algorithm}
                        aria-label="Sorting algorithm"
                        onChange={(event) => {
                            setAlgorithm(event.target.value);
                            resetRun();
                            setStatus("");
                        }}
                    >
                        {ALGORITHMS.map((item) => (
                            <option key={item.value} value={item.value}>
                                {item.label}
                            </option>
                        ))}
                    </select>
                </div>

                <button className="control-button" onClick={playing ? () => setPlaying(false) : start}>
                    <i className={playing ? "fas fa-pause" : "fas fa-play"}></i>{" "}
                    {playing ? "Stop" : "Start"}
                </button>

                <button className="control-button" onClick={shuffle}>
                    <i className="fas fa-random"></i> Shuffle
                </button>

                <div className="speed-control">
                    <label htmlFor="speed">
                        <i className="fa-solid fa-gauge"></i> Speed:
                    </label>
                    <input
                        id="speed"
                        type="range"
                        min="2"
                        max="11"
                        value={speedValue}
                        onChange={(event) => setSpeedValue(Number(event.target.value))}
                    />
                </div>

                <button className="control-button" onClick={showComplexity}>
                    <i className="fa-solid fa-chart-line"></i> Complexity
                </button>

                <button className="control-button" onClick={showCode}>
                    <i className="fa-solid fa-code"></i> Code
                </button>

                <p className="status-text" role="status">
                    {status}
                </p>
            </aside>

            <main className="main-content">
                <h1>🔮SortViz</h1>
                <div id="visualizer" aria-label="Sorting visualization">
                    {displayBars.map((value, index) => (
                        <div
                            className="bar"
                            key={index}
                            style={{ height: `${value * 9.5}px` }}
                        >
                            {value}
                        </div>
                    ))}
                </div>
                <p className="vis-caption">
                    Step {steps ? Math.min(stepIndex + 1, steps.length) : 0} of{" "}
                    {steps ? steps.length : 0}
                </p>
            </main>

            {modal && (
                <Modal title={modal.title} onClose={() => setModal(null)}>
                    {modal.code ? <pre>{modal.code}</pre> : <p>{modal.text}</p>}
                </Modal>
            )}
        </div>
    );
}

function AlgorithmsPage() {
    const [data, setData] = useState(null);
    const [error, setError] = useState("");

    useEffect(() => {
        fetch("/static/data/algorithm_data.json")
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Could not load algorithm data");
                }
                return response.json();
            })
            .then(setData)
            .catch((err) => setError(err.message));
    }, []);

    return (
        <main className="page-content">
            <h1>Sorting Algorithms</h1>
            <p className="page-lead">
                Compare the complexity and reference implementation of every algorithm
                available in the visualizer.
            </p>

            {error && <p className="error-text">{error}</p>}
            {!data && !error && <p className="loading-text">Loading algorithms…</p>}

            <div className="algorithm-grid">
                {data &&
                    ALGORITHMS.map((item) => {
                        const info = data[item.value];
                        return (
                            <article className="algorithm-card" key={item.value}>
                                <h2>{item.label}</h2>
                                <p className="complexity">
                                    {info ? info.complexity : "Complexity unavailable"}
                                </p>
                                {info && (
                                    <pre>
                                        {info.code.replace(/\{\{NEWLINE\}\}/g, "\n")}
                                    </pre>
                                )}
                            </article>
                        );
                    })}
            </div>
        </main>
    );
}

function AboutPage() {
    return (
        <main className="page-content">
            <h1>About SortViz</h1>
            <section className="about-card">
                <p>
                    SortViz is a sorting algorithm visualizer. It replays each swap the
                    backend records so you can watch an array move toward its sorted
                    order, one step at a time.
                </p>
                <p>
                    The server is written in Go and exposes a single <code>POST /sort</code>{" "}
                    endpoint. This frontend is built with React and split into separate
                    pages for the visualizer, the algorithm reference and this overview.
                </p>
                <p>
                    Eight algorithms are available: selection, bubble, insertion, merge,
                    quick, heap, shell and cocktail shaker sort. Use the Controls panel to
                    pick one, change the animation speed, or inspect its complexity and
                    reference implementation.
                </p>
                <p>
                    <a className="text-link" href="/">
                        Open the visualizer
                    </a>
                </p>
            </section>
        </main>
    );
}

function App() {
    const page = document.body.dataset.page;

    let content = <Visualizer />;
    if (page === "algorithms") {
        content = <AlgorithmsPage />;
    } else if (page === "about") {
        content = <AboutPage />;
    }

    return (
        <React.Fragment>
            <Nav />
            {content}
        </React.Fragment>
    );
}

ReactDOM.createRoot(document.getElementById("root")).render(<App />);
