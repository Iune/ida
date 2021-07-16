import fs from 'fs';

export function loadJSON(path: string): object {
    return JSON.parse(fs.readFileSync(path, 'utf-8'));
}

export function writeJSON(json: object, path: string) {
    fs.writeFileSync(path, JSON.stringify(json));
}

export function loadTextLines(path: string): string[] {
    return fs.readFileSync(path, 'utf-8')
        .split('\n')
        .map(line => line.trim());
}

// https://stackoverflow.com/a/46700791
export function notEmpty<T>(value: T | null | undefined): value is T {
    return value !== null && value !== undefined;
}