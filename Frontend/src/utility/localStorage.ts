



export function addToLocalStorage(key: string, value: any): void {
    window.localStorage.setItem(key, JSON.stringify(value))
}

export function getValueFromLocalStorage(key: string): any {
    return JSON.parse(<string>window.localStorage.getItem(key))
}