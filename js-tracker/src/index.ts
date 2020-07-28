const ERROR_PROCESSOR_URL = "http://localhost:8088";

(() => {
  function sendError(
    msg: string | Event,
    url: string | undefined,
    lineNo: number | undefined,
    columnNo: number | undefined,
    err: Error | undefined
  ) {
    fetch(ERROR_PROCESSOR_URL + "/receive-error", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        msg,
        url,
        lineNo: lineNo?.toString(),
        columnNo: columnNo?.toString(),
        err: err?.name,
        website: window.location.origin,
        teamId: "test-123",
      }),
    });
  }

  window.onerror = function (
    msg: string | Event,
    url: string | undefined,
    lineNo: number | undefined,
    columnNo: number | undefined,
    err: Error | undefined
  ): boolean {
    try {
      sendError(msg, url, lineNo, columnNo, err);
    } catch (e) {
      console.log(e);
    }
    return false;
  };
})();
