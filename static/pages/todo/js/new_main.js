class App {
  constructor() {
    this.matrix = {
      list: { items: {} },
      item: { subItems: {}, complete: false, text: "", id: null },
      subItem: { complete: false, text: "", id: null },
      user: { browser: null, entropy: null, id: null, settings: null },
      api: {
        get: {
          route: "/",
          headers: {
            "X-Action": "GET",
            Accept: "application/json",
            "Content-Type": "application/json",
          },
          method: "POST",
        },
        set: {
          route: "/",
          headers: {
            "X-Action": "SET",
            Accept: "application/json",
            "Content-Type": "application/json",
          },
          method: "POST",
        },
        del: {
          route: "/",
          headers: {
            "X-Action": "DEL",
            Accept: "application/json",
            "Content-Type": "application/json",
          },
          method: "POST",
        },
      },
    };
    this.schema = {
      placeholders: {
        add_subitem: "Add subitem",
        add_item: "Add item",
      },
      settings: {
        storage_method: "cloud", //Options are cloud, local and none
        color_scheme: "dark", //Options are light and dark
      },
      css_assets: {
        remove_icon: `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash2" viewBox="0 0 16 16"><path d="M14 3a.702.702 0 0 1-.037.225l-1.684 10.104A2 2 0 0 1 10.305 15H5.694a2 2 0 0 1-1.973-1.671L2.037 3.225A.703.703 0 0 1 2 3c0-1.105 2.686-2 6-2s6 .895 6 2zM3.215 4.207l1.493 8.957a1 1 0 0 0 .986.836h4.612a1 1 0 0 0 .986-.836l1.493-8.957C11.69 4.689 9.954 5 8 5c-1.954 0-3.69-.311-4.785-.793z"/></svg>`,
        add_item_icon: `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-earmark-plus" viewBox="0 0 16 16"><path d="M8 6.5a.5.5 0 0 1 .5.5v1.5H10a.5.5 0 0 1 0 1H8.5V11a.5.5 0 0 1-1 0V9.5H6a.5.5 0 0 1 0-1h1.5V7a.5.5 0 0 1 .5-.5z"/><path d="M14 4.5V14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h5.5L14 4.5zm-3 0A1.5 1.5 0 0 1 9.5 3V1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4.5h-2z"/></svg`,
        checkmark_icon: `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-check-circle" viewBox="0 0 16 16"><path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/><path d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"/></svg>`,
      },
    };
    this.state = {
      list: Object.assign({}, this.schema.list),
      user: Object.assign({}, this.schema.user),
    };
    this.d = document;
    return this;
  }

  init() {
    this.d.body.addEventListener("click", this.keyboardHook);
    this.d.body.addEventListener("keydown", this.mouseHook);
  }

  keyboardHook(e) {}

  mouseHook(e) {}

  async setUser() {
    this.state.user.browser = this.getBrowserInfo();
    this.state.user.entropy = this.getFingerPrint();
    this.state.user.id = await hash(
      this.state.user.entropy + this.state.user.browser,
    );
    return this;
  }

  async fetchState() {
    // TODO: implement local and nil storage
    const r = await fetch(
      `${this.matrix.api.get.route}?user_id=${this.state.user.id}`,
      {
        method: this.matrix.api.get.method,
        headers: this.matrix.api.get.headers,
        body: "",
      },
    );
    const ro = await r.json();
    if (ro?.data && this.validateState(ro.data)) {
      this.state.list = ro.data;
    }
    return this;
  }

  async pushState() {
    // TODO:implement local and nil storage
    await fetch(`${this.matrix.api.set.route}?user_id=${this.state.user.id}`, {
      method: this.matrix.api.set.method,
      headers: this.matrix.api.set.headers,
      body: JSON.stringify(this.state.list),
    });
    return this;
  }

  validateState(object) {
    if (!object?.items) {
      return false;
    }
    if (!object.items[0]?.id) {
      return false;
    }
    return true;
  }

  getFingerPrint() {
    const c = this.d.createElement(canvas);
    c.id = "fingerprint";
    c.width = 300;
    c.height = 300;
    this.d.body.appendChild(c);
    const ctx = c.getContext("2d");
    const txt = "sTub>BbeE,ioN#$|`` <cØan.vas>:` 24601il";
    ctx.textBaseline = "top";
    ctx.font = "14px 'Arial'";
    ctx.textBaseline = "alphabetic";
    ctx.fillStyle = "#f60";
    ctx.fillRect(125, 1, 62, 20);
    ctx.fillStyle = "#069";
    ctx.fillText(txt, 2, 15);
    ctx.fillStyle = "rgba(102, 204, 0, 0.7)";
    ctx.fillText(txt, 4, 17);
    const hsh = c.toDataURL();
    c.remove();
    return hsh.substring(hsh.length - 128);
  }

  getBrowserInfo() {
    const wn = window.navigator,
      ws = window.screen;
    return `${wn.deviceMemory ?? 0}${wn.hardwareConcurrency}${wn.language}${
      wn.maxTouchPoints
    }${ws.height + ws.width + ws.pixelDepth}`;
  }

  async hash(string) {
    const hashBuffer = await crypto.subtle.digest(
      "SHA-256",
      new TextEncoder().encode(string),
    );
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    return hashArray
      .map((bytes) => bytes.toString(16).padStart(2, "0"))
      .join("");
  }
}

new App().init();
