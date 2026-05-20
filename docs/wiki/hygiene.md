# Hygiene and Git Workflow

This project follows the **Jules Dev Standard**.

## Atomic Commits

Conventional Commits format:
`<type>(<scope>): <subject>`

### Allowed Types

- `feat`: New functionality.
- `fix`: Bug correction.
- `docs`: Documentation changes.
- `style`: Visual changes (no logic).
- `refactor`: Code change that neither adds nor fixes anything.
- `chore`: Maintenance tasks, dependencies.

## Branch Workflow

- `main`: Production branch (linear history only).
- `feat/*`: Branches for new functionalities.
- `fix/*`: Branches for corrections.

**Banned:** `git push --force` to `main`.
