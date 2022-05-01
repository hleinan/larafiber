var __defProp = Object.defineProperty;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __publicField = (obj, key, value) => {
  __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
  return value;
};
import { LitElement, css, html, render } from "https://cdn.skypack.dev/lit-element";
class ArtPlants extends LitElement {
  static get styles() {
    return css`
      :host {
        display: block;
        color: var(--art-plants-text-color, #000);
        font-family: neusa-next-std,sans-serif;
      }

      h1, h2, h3, h4, h5, h6 {
        font-family: korolev-compressed,sans-serif !important;
      }

      .artplant-ul {
        list-style: none;
        padding: 0;
      }
      
      a {
        color: #212121;
        text-decoration: none;

      }
      .artplant-link {
        display: flex;
        margin-bottom: 10px;
        background: #ffffff;
        padding: 10px;
      }

      .artplant-link:hover {
        color: #0077b3;
      }

      .artplant-li-month {
        text-transform: uppercase;
      }
      
      .artplant-li-day {
        margin: 4px 0;
        font-size: 3em;
        font-family: korolev-compressed,sans-serif !important;
      }

      .artplant-date-box {
        margin-right: 20px;
        text-align: center;
        padding: 5px 20px;
        background: #f5f5f5;
        border-radius: 4px;
        line-height: 1.8em;
      }
      .artplant-description {
        font-size: 2.3em;
        margin-bottom: 2px;
        line-height: 1.5em;
        font-family: korolev-compressed,sans-serif !important;
      }
      .artplant-date {
        font-size: 0.8em;
        line-height: 1.25em;
      }
      
      .artplant-location {
        font-size: 1em;
        line-height: 1.25em;
      }
    `;
  }
  constructor() {
    super();
  }
  connectedCallback() {
    super.connectedCallback();
    fetch("https://app.artplant.app/api/v1/artplant/" + this.type + "/" + this.hash).then((response) => response.json()).then((json) => this._events = json);
  }
  render() {
    var _a;
    return html`
      <ul class="artplant-ul">
        ${(_a = this._events) == null ? void 0 : _a.map((event) => html`
          <li class="artplant-li">
            <a class="artplant-link" href="${event.url}" target="_blank">
              <div class="artplant-date-box">
                <div>${new Date(event.start).toLocaleString('no-NO', { weekday: 'short' })}</div>
                <div class="artplant-li-day stylish">${new Date(event.start).getDate()}</div>
                <div class="artplant-li-month">${new Date(event.start).toLocaleString('no-NO', { month: 'short' })}</div>
              </div>
              <div>
                <div class="stylish artplant-description">${event.name}</div>
                <div class="artplant-date">${new Date(event.start).toLocaleString('no-NO', {day:'2-digit', month:'2-digit', year:'numeric', hour:'2-digit', minute:'2-digit'})}</div>
                <div class="artplant-location">${event.venue.name} | ${event.venue.address.city}</div>
              </div>
            </a>
          </li>
          `)}
      </ul>
    `;
  }
}
__publicField(ArtPlants, "properties", {
  hash: { type: String },
  type: { type: String },
  _events: { type: Array }
});
window.customElements.define("art-plants", ArtPlants);
const ArtplantEvents = (hash, type) => {
  render(html`
        <art-plants .hash=${hash} .type=${type}>
          some light-dom
        </art-plants>
      `, document.querySelector("#artplant"));
};
export { ArtplantEvents as default };
