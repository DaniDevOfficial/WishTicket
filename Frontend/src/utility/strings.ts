export function truncation(stringToTruncate: string, maxLength: number = 100, trailingValue: string = "") {
    if(stringToTruncate.length < maxLength) return stringToTruncate;

    const truncatedString = stringToTruncate.substring(0, maxLength)

    return truncatedString + trailingValue

}