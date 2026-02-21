type DodoPaymentsAPIError = {
    error: {
        code: 'NOT_FOUND';
        message: string;
    }
}

export function isDodoPaymentsAPIError(e: unknown): e is DodoPaymentsAPIError {
    return typeof e === 'object' && e !== null && 'error' in e && typeof (e as any).error?.code === 'string';
}