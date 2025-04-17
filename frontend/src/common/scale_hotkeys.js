export function IS_HOTKEY_UP_FUNCTION(event) { 
    return event.ctrlKey && (event.key === "+" || event.key === "=");
  }
  
  export function IS_HOTKEY_DOWN_FUNCTION(event) {
    return event.ctrlKey && event.key === "-";
  }