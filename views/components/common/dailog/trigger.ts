function openModal(e: CustomEvent) {
  const payload = JSON.parse(e.detail.elt.dataset.action);
  if (!payload) {
    throw new Error("Payload is required for modal");
  }
  renderDialog(payload);
  handleDialogActions(e);
}

window.openModal = openModal;

const renderDialog = (payload: DialogProps) => {
  const dialog = document.createElement("dialog");
  dialog.id = "my-modal";
  dialog.classList.add("modal", "modal-bottom", "sm:modal-middle");
  dialog.innerHTML = `
  <div class="modal-box">
    <h3 class="font-bold text-lg" id="modal-title">${payload.title}</h3>
    <p class="py-4" id="modal-description">${payload.description}</p>
    <div class="modal-action">
      <button type="button" class="btn" id="modal-close-btn">Close</button>
      <button type="button" class="btn btn-primary" id="modal-confirm-btn">Confirm</button>
    </div>
  </div>
  `;
  document.body.insertAdjacentElement("beforeend", dialog);

  const modal = document.getElementById("my-modal") as HTMLDialogElement;
  modal.showModal();
};

const handleDialogActions = (e: CustomEvent) => {
  const modal = document.getElementById("my-modal") as HTMLDialogElement;
  const closeBtn = document.getElementById("modal-close-btn");
  const confirmBtn = document.getElementById("modal-confirm-btn");

  const removeModal = (modal: HTMLDialogElement) => {
    modal.remove();
  };

  closeBtn.addEventListener("click", () => {
    removeModal(modal);
  });

  confirmBtn.addEventListener("click", () => {
    e.detail.issueRequest();
    modal.remove();
  });
};

type DialogProps = {
  title: string;
  description: string;
  id: string;
};
