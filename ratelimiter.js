class RateLimiter {
  constructor(limit, windowMs) {
    this.limit = limit;
    this.windowMs = windowMs;
    this.requests = new Map();
  }

  isAllowed(ip) {
    const now = Date.now();

    if (!this.requests.has(ip)) {
      this.requests.set(ip, []);
    }

    const timestamps = this.requests.get(ip)
      .filter(ts => now - ts < this.windowMs);

    timestamps.push(now);

    this.requests.set(ip, timestamps);

    return timestamps.length <= this.limit;
  }
}
