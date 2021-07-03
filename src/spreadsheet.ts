import { Entry } from "./contest";
import { CellFormulaValue, CellValue, Workbook, Worksheet, WorksheetView } from 'exceljs';

export function createWorkbook(entries: Entry[]): Workbook {
    // Create workbook
    const workbook = new Workbook();
    const sheet = workbook.addWorksheet();
    // Add rows
    sheet.addRow(getHeaderRowContents(entries));
    entries.forEach((entry, index) => sheet.addRow(getEntryRowContents(entry, index + 1)));
    // Set formulas
    setFormulas(sheet);
    // Set voter placeholders ('X')
    setVoterPlaceholders(sheet);
    // Set cell formatting
    setFormats(sheet);
    return workbook;
}

function getHeaderRowContents(entries: Entry[]): string[] {
    const headerRow = ['Draw', 'Country', 'Flag', 'Artist', 'Song', 'Total', 'Count']
    entries.forEach(entry => headerRow.push(entry.country.primaryName() ?? ''));
    return headerRow;
}

function getEntryRowContents(entry: Entry, draw: number): Array<string | number> {
    return [draw, entry.country.primaryName() ?? '', entry.flag(), entry.artist, entry.song];
}

function setFormulas(sheet: Worksheet) {
    for (let i = 2; i <= sheet.rowCount; i++) {
        const voteFirstColCellAddress = sheet.getCell(i, 8).address;
        const voteLastColCellAddress = sheet.getCell(i, sheet.columnCount - 1).address;
        sheet.getCell(i, 6).value = { formula: `SUM(${voteFirstColCellAddress}:${voteLastColCellAddress})`, result: 0 } as CellFormulaValue;
        sheet.getCell(i, 7).value = { formula: `COUNT(${voteFirstColCellAddress}:${voteLastColCellAddress})`, result: 0 } as CellFormulaValue;
    }
}

function setVoterPlaceholders(sheet: Worksheet) {
    for (let i = 2; i <= sheet.rowCount; i++) {
        sheet.getCell(i, i + 6).value = 'X';
    }
}

function setFormats(sheet: Worksheet) {
    // Set font details and text-alignment for header row
    sheet.getRow(1).font = { name: 'Palatino', family: 1, size: 12, bold: true };
    sheet.getRow(1).alignment = { horizontal: 'center', vertical: 'bottom' };
    // Set font details and text-alignment for remaining rows
    for (let i = 2; i <= sheet.rowCount; i++) {
        sheet.getRow(i).font = { name: 'Palatino', family: 1, size: 12 };
        sheet.getRow(i).alignment = { horizontal: 'center', vertical: 'middle' };
    }
    // Rotate total, count and voter name columns in header row
    for (let i = 6; i <= sheet.columnCount; i++) {
        sheet.getCell(1, i).alignment = { textRotation: 90, horizontal: 'center', vertical: 'bottom' };
    }
}

export async function writeWorkbook(workbook: Workbook, path: string) {
    await workbook.xlsx.writeFile(path);
}