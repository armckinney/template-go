---
description: Wipe devcontainer without cache by removing containers and images
---

// turbo-all
1. Stop the dev container, remove volumes, and remove associated images
```bash
docker compose -p template-go_devcontainer -f .devcontainer/docker-compose.yaml down --volumes --rmi all
```

2. Run the Antigravity command to reopen in container
> **Action Required**: Open the Command Palette (`Cmd+Shift+P`) and run **"Remote-Containers: Reopen Folder in Container"** (or click the "Reopen in Container" button in the UI).