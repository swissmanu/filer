<script lang="ts">
  import pdfjs from "pdfjs-dist";
  import pdfjsWorker from "pdfjs-dist/build/pdf.worker.entry";
  import { onMount } from "svelte";

  export let url: string;

  pdfjs.GlobalWorkerOptions.workerSrc = pdfjsWorker;
  let canvasRef: HTMLCanvasElement;
  let measureRef: HTMLDivElement;
  let canvasWidth = 0;

  let pdfDocument: pdfjs.PDFDocumentProxy;
  let currentPageNumber = 1;
  let pending = false;

  $: {
    loadDocument(url);
  }

  $: if (pdfDocument !== undefined) {
    renderPage(currentPageNumber, canvasWidth);
  }

  const loadDocument = async (url: string) => {
    pending = true;
    const document = await pdfjs.getDocument(url).promise;
    pdfDocument = document;
    currentPageNumber = 1;
  };

  const renderPage = async (pageNumber: number, availableWidth: number) => {
    pending = true;

    const page = await pdfDocument.getPage(pageNumber);
    const viewport = page.getViewport({ scale: 1 });
    // canvasRef.width = viewport.width;
    canvasRef.height = viewport.height;
    const scaledViewport = page.getViewport({
      scale: availableWidth / viewport.width,
    });
    measureRef.style.height = `${scaledViewport.height}px`;
    page.render({
      canvasContext: canvasRef.getContext("2d"),
      viewport: scaledViewport,
    });

    pending = false;
  };

  onMount(() => {
    const observer = new ResizeObserver(
      ([
        {
          contentRect: { width },
        },
      ]) => (canvasWidth = width)
    );
    observer.observe(measureRef);
  });
</script>

<div class="p-8 shadow-inner bg-gray-50 rounded">
  <div
    class="relative border border-gray-200 overflow-hidden"
    bind:this={measureRef}
  >
    <canvas
      class="absolute"
      width={canvasWidth}
      height={400}
      bind:this={canvasRef}
    />
    {#if pending}
      <div>Loading...</div>
    {/if}
  </div>
</div>

<style>
</style>
