const API_BASE_URL = "http://localhost:3001/api/v1"

class ErrorMetricsService {
  private async sendError(error: Error, context: string) {
    try {
      const response = await fetch(`${API_BASE_URL}/metrics/error`, {
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

      if (!response.ok) {
        console.error("Error response:", await response.text())
      }
    } catch (error_) {
      console.error("Failed to send error metric:", error_)
    }
  }

  public reportError(error: Error, context: string) {
    this.sendError(error, context).catch(console.error)
  }
}

export const errorMetrics = new ErrorMetricsService()