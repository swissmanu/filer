<script lang="ts">
  import pdfjs from "pdfjs-dist";
  import pdfjsWorker from "pdfjs-dist/build/pdf.worker.entry";

  export let url: string;

  pdfjs.GlobalWorkerOptions.workerSrc = pdfjsWorker;
  let canvasRef: HTMLCanvasElement;

  let pdfDocument: pdfjs.PDFDocumentProxy;
  let currentPageNumber = 1;
  let pending = false;

  $: {
    loadDocument(url);
  }

  $: if (pdfDocument !== undefined) {
    renderPage(currentPageNumber);
  }

  const loadDocument = async (url: string) => {
    pending = true;
    const document = await pdfjs.getDocument(url).promise;
    pdfDocument = document;
    currentPageNumber = 1;
  };

  const renderPage = async (pageNumber: number) => {
    pending = true;
    const page = await pdfDocument.getPage(pageNumber);
    const viewport = page.getViewport({ scale: 1, rotation: 0 });
    canvasRef.width = viewport.width;
    canvasRef.height = viewport.height;
    page.render({ canvasContext: canvasRef.getContext("2d"), viewport });
    pending = false;
  };
</script>

<style>
</style>

<div>
  <canvas width={300} height={400} bind:this={canvasRef} />
  {#if pending}
    <div>Loading...</div>
  {/if}
</div>
