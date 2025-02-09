package capture

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
	"runtime"
)

func Screenshot(filepath string) (image.Image, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	htmlContent := generateHTML(string(content))
	tmpFile, err := os.CreateTemp("", "code-*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	if err := os.WriteFile(tmpFile.Name(), []byte(htmlContent), 0644); err != nil {
		return nil, fmt.Errorf("failed to write temp file: %w", err)
	}

	var cmd *exec.Cmd
	chrome := "google-chrome"
	if runtime.GOOS == "darwin" {
		chrome = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	}

	outputFile := filepath + ".png"
	cmd = exec.Command(
		chrome,
		"--headless",
		"--disable-gpu",
		"--screenshot="+outputFile,
		"--window-size=1280,800",
		"file://"+tmpFile.Name(),
	)

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to capture screenshot: %w", err)
	}

	imgFile, err := os.Open(outputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open screenshot: %w", err)
	}
	defer imgFile.Close()
	defer os.Remove(outputFile)

	img, err := png.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode screenshot: %w", err)
	}

	return img, nil
}

func generateHTML(content string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/github.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/highlight.min.js"></script>
    <style>
        body { 
            background: white; 
            padding: 20px; 
            margin: 0;
            min-width: 800px;
        }
        pre { 
            margin: 0; 
            background: #ffffff;
            padding: 15px;
        }
        code { 
            font-family: 'Monaco', 'Consolas', monospace; 
            font-size: 14px; 
            line-height: 1.4;
        }
    </style>
</head>
<body>
    <pre><code>%s</code></pre>
    <script>
        document.addEventListener('DOMContentLoaded', (event) => {
            document.querySelectorAll('pre code').forEach((el) => {
                hljs.highlightElement(el);
            });
        });
    </script>
</body>
</html>
`, content)
}
