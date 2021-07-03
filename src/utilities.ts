import fs from 'fs';

export function loadJSON(path: string): object {
    return JSON.parse(fs.readFileSync(path, 'utf-8'));
}

export function loadTextLines(path: string): string[] {
    return fs.readFileSync(path, 'utf-8')
        .split('\n')
        .map(line => line.trim());
}