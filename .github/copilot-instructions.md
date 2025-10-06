# Radicaster

Radicaster is a Go application for recording Japanese radio content from Radiko and serving it as podcast feeds. The application runs as an HTTP server that manages scheduled recordings and serves podcast RSS feeds.

**Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

## Working Effectively

### Bootstrap and Build (FAST - No timeout needed)
- **CRITICAL**: Ensure Go 1.21+ is installed: `go version`
- Install required external dependencies: `sudo apt-get update && sudo apt-get install -y ffmpeg`
- Download dependencies: `go mod download` -- takes 10-30 seconds
- Build the application: `make build` -- takes <1 second. NEVER CANCEL.
- Alternative build: `go build -o dist/radicaster cmd/radicaster/main.go`

### Testing (FAST - No timeout needed)
- Run working tests: `go test ./config/ ./timeutil/` -- takes <1 second. NEVER CANCEL.
- **CRITICAL**: `go test ./record/` and `go test ./...` will FAIL due to network dependencies (radiko.jp). This is expected in isolated environments.
- Ensure testdata exists: The file `config/testdata/radicaster.yaml` is required for config tests.
- Format check: `gofmt -l .` -- takes <1 second
- Module validation: `go mod verify && go mod tidy` -- takes 5-10 seconds

### Development Environment
- Hot reload development: `export PATH=$PATH:~/go/bin && go install github.com/air-verse/air@latest && make watch`
- **CRITICAL**: `make watch` will start and immediately crash due to network dependencies. This is expected.
- Configuration file: Create `radicaster.yaml` in the root directory for testing

### Running the Application
- **CRITICAL**: The application requires internet access to radiko.jp and will panic in isolated environments
- Basic usage: `./dist/radicaster -baseurl http://localhost:3333 -config ./radicaster.yaml -targetdir ./output`
- With authentication: `./dist/radicaster -baseurl http://localhost:3333 -config ./radicaster.yaml -targetdir ./output -basicauth user:password`
- Premium features: Add `-radikoemail "${RADIKO_EMAIL}" -radikopassword "${RADIKO_PASSWORD}"`
- **CRITICAL**: Always set `ulimit -n 16384` before running in production

## Validation

### Manual Testing
- **CRITICAL**: The application cannot be fully tested without internet access to radiko.jp
- Always test configuration parsing: `go test ./config/`
- Always test time utilities: `go test ./timeutil/`
- **VALIDATION SCENARIO**: After making changes, verify the build succeeds and configuration parsing works
- Check that configuration files validate: Create a test YAML file and ensure the app accepts it

### Required External Dependencies
- **ffmpeg**: Required for audio processing. Install with `sudo apt-get install -y ffmpeg`
- **Internet access**: Required for radiko.jp API access
- **File system permissions**: Application needs read/write access to target directory

### Pre-commit Validation
- Always run: `gofmt -l .` (should return no output)
- Always run: `go mod tidy`
- Always run: `go test ./config/ ./timeutil/`
- Always run: `make build` to ensure compilation succeeds

## Common Tasks

### Repository Structure
```
/home/runner/work/radicaster/radicaster/
├── cmd/radicaster/main.go           # Main application entry point
├── config/                          # Configuration parsing and validation
├── record/                          # Recording logic (network dependent)
├── podcast/                         # Podcast RSS feed generation
├── http/                            # HTTP server and web interface
├── ffmpeg/                          # Audio processing utilities
├── radikoutil/                      # Radiko API client utilities
├── timeutil/                        # Time and scheduling utilities
├── metadata/                        # Audio metadata handling
├── Makefile                         # Build commands (build, watch, clean)
├── .air.toml                        # Hot reload configuration
├── go.mod                           # Go module definition
└── README.md                        # Basic usage documentation
```

### Key Configuration Files
- `radicaster.yaml`: Main configuration file (see README.md for format)
- `.air.toml`: Development hot reload configuration
- `.goreleaser.yaml`: Release build configuration
- `config/testdata/radicaster.yaml`: Test configuration (required for tests)

### Common Commands Output
```bash
# Repository root contents
ls -la /home/runner/work/radicaster/radicaster/
.air.toml  .github/  .gitignore  .goreleaser.yaml  Makefile  README.md  cmd/  config/  dist/  ffmpeg/  go.mod  go.sum  http/  metadata/  output/  podcast/  radikoutil/  record/  timeutil/

# Build output
make build
# Creates: dist/radicaster (binary ~21MB)

# Working tests
go test ./config/ ./timeutil/
# PASS: config (0.003s), timeutil (0.005s)

# Module verification
go mod verify
# all modules verified
```

### Network Dependencies and Limitations
- **CRITICAL**: Application requires access to radiko.jp for radio content
- **CRITICAL**: Tests in `./record/` package will fail in isolated environments
- **CRITICAL**: Running the application without network access will cause panic
- The application is designed for deployment in environments with full internet access
- Use configuration files to test parsing logic without network requirements

### Development Workflow
1. Always start with: `go mod download && make build`
2. Test configuration changes: `go test ./config/`
3. For new time logic: `go test ./timeutil/`
4. Format code: `gofmt -w .`
5. Verify build: `make build`
6. **NEVER** expect full application testing without internet access

### Performance Characteristics
- Build time: <1 second
- Test time (working tests): <1 second
- Module download: 10-30 seconds
- Binary size: ~21MB
- **CRITICAL**: No long-running operations in build/test cycle - all timeouts can be minimal

### Error Patterns to Expect
- Network errors: `dial tcp: lookup radiko.jp` - expected in isolated environments
- Configuration errors: Check YAML syntax and required fields
- FFmpeg errors: Ensure ffmpeg is installed and in PATH
- File permission errors: Ensure target directory is writable

### Essential Files for Testing
Always ensure these files exist:
- `config/testdata/radicaster.yaml` (created automatically by test setup)
- `radicaster.yaml` (create manually for testing configuration parsing)
- `output/` directory (created automatically by application)