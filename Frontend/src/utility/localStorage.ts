
export function addToLocalStorage(key: string, value: any): void {
    try {
        const serializedValue = JSON.stringify(value);
        window.localStorage.setItem(key, serializedValue);
    } catch (error) {
        console.error(`Error saving to localStorage: ${error}`);
    }
}

export function getValueFromLocalStorage(key: string): any | null {
    try {
        const storedValue = window.localStorage.getItem(key);

        if (!storedValue) {
            return null;
        }

        return JSON.parse(storedValue);
    } catch (error) {
        // console.error(`Error retrieving or parsing from localStorage: ${error}`);
        return null;
    }
}
