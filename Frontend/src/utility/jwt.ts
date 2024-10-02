
export interface jwtBody {
    Exp: bigint,
    UserId: number,
    UserName: string,
}

export function getJWTBody(jwt: string): jwtBody | null {
    if (!jwt) {
        throw new Error('No JWT');
    }
    const parts = jwt.split('.');

    if (parts.length !== 3) {
        throw new Error('Invalid JWT format');
    }

    const payload = parts[1];

    return JSON.parse(decodeBase64Url(payload));
}

function decodeBase64Url(base64Url: string) {
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');

    const padding = base64.length % 4 === 0 ? '' : '='.repeat(4 - (base64.length % 4));

    return decodeURIComponent(escape(window.atob(base64 + padding)));
}


