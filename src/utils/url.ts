export const HTTP_SCHEME_ERROR = 'URL must start with http:// or https://';

export function isHttpUrl(value: string): boolean {
  return value.startsWith('http://') || value.startsWith('https://');
}
