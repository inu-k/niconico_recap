export async function fetchData(url: string, options: any = {}) {
    const response = await fetch(url, options);
    if (!response.ok) {
        throw new Error('Failed to fetch data');
    }
    const data = await response.json();
    return data;

}

export function formatDate(date: string) {
    const createdAt = new Date(date);

    const year = createdAt.getFullYear();
    const month = ('0' + (createdAt.getMonth() + 1)).slice(-2);
    const day = ('0' + createdAt.getDate()).slice(-2);
    const hours = ('0' + createdAt.getHours()).slice(-2);
    const minutes = ('0' + createdAt.getMinutes()).slice(-2);
    const seconds = ('0' + createdAt.getSeconds()).slice(-2);

    const formattedData = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return formattedData;
}