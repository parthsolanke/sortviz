# ![Icon](./web/static/assets/favicon.ico) SortViz

SortViz is a Go web application that visualizes sorting algorithms step by step in the browser.

![SortViz Screenshot](./web/static/assets/app.png)

### Features
- Visualizes each step of the sorting process as animated bars.
- Supports Selection, Bubble, Insertion, Merge and Quick sort.
- Shuffle, start/pause and adjustable animation speed controls.
- Per-algorithm complexity and reference code popups.
- Small, dependency-free Go backend using only the standard library.

### Project layout
- `cmd/server/main.go` — HTTP entrypoint that wires up the router and serves the app on port 8080.
- `internal/handler/` — `IndexHandler` renders the page; `SortHandler` implements the sort API.
- `internal/sorter/` — sorting implementations that record intermediate `SortStep` values.
- `web/templates/index.html` — the page template.
- `web/static/` — CSS, JavaScript, images and `data/algorithm_data.json`.
- `Makefile`, `Dockerfile`, `.air.toml` — build, container and live-reload helpers.

### Prerequisites
- Go 1.22.4 or later (matches `go.mod`)
- Make (optional, for the Makefile targets)
- Docker (optional, for running in a container)

### Running locally
```bash
go run cmd/server/main.go
```
Then open `http://localhost:8080`. The server reads templates and static files from `web/`, so run it from the repository root.

### Using the Makefile
```bash
make build   # builds bin/sortviz.exe from cmd/server/main.go
make test    # runs go test -v ./internal/...
make run     # builds and runs the binary
make clean   # removes the bin/ directory
make         # clean, build, test, run
```

### Docker
```bash
docker build -t sortviz .
docker run --rm -p 8080:8080 sortviz
```
The image builds the binary with the Makefile and copies `web/` into the runtime image.

### Routes and API
- `GET /` — renders `web/templates/index.html`.
- `GET /static/...` — serves files from `web/static`.
- `POST /sort` — accepts JSON and returns the recorded steps.

Request body:
```json
{ "array": [5, 3, 8, 1], "algorithm": "quick" }
```

`algorithm` must be one of `bubble`, `insertion`, `selection`, `merge` or `quick`; anything else returns `400 Unknown sorting algorithm`. Invalid JSON returns `400`.

Response (a `[]sorter.SortStep`, one entry per recorded array state):
```json
[
  { "array": [3, 5, 1, 8] },
  { "array": [1, 3, 5, 8] }
]
```

### Usage
Pick an algorithm from the sidebar, press Shuffle to randomize the bars, then Start to animate the sort; the button toggles to Stop while running. Use the speed slider to slow down or speed up the animation, and the Complexity and Code buttons to show details for the selected algorithm.

### License
This project is licensed under the [MIT License](./LICENSE).
