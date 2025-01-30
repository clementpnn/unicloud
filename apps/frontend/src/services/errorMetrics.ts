class ErrorMetricsService {
  private async sendError(error: Error, context: string) {
    try {
      await fetch("http://localhost:3000/api/v1/metrics/error", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          error_type: error.name,
          error_message: error.message,
          context: context
        })
      })
    } catch (error_) {
      console.error("Failed to send error metric:", error_)
    }
  }

  public reportError(error: Error, context: string) {
    this.sendError(error, context)
  }
}

export const errorMetrics = new ErrorMetricsService()