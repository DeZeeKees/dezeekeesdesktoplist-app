import Swal from "sweetalert2";

export const Toast = Swal.mixin({
    toast: true,
    position: "top-end",
    showConfirmButton: false,
    timer: 3000,
    customClass: {
        popup: "toast-down"
    },
    didOpen: (toast) => {
      toast.onmouseenter = Swal.stopTimer;
      toast.onmouseleave = Swal.resumeTimer;
    }
});

export function capitalize(string) {
    // replace all underscores with spaces
    string = string.replace(/_/g, " ");
    // capitalize first letter
    return string.charAt(0).toUpperCase() + string.slice(1);
}

export function formatNumber(num) {
    return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

export function formatRating(rating) {
    switch(rating) {
        case "g":
            return "G - All Ages";
        case "pg":
            return "PG - Children";
        case "pg_13":
            return "PG-13 - Teens 13 or older";
        case "r":
            return "R - 17+ (violence & profanity)";
        case "r+":
            return "R+ - Mild Nudity";
        case "rx":
            return "Rx - Hentai";
        default:
            return "N/A";
    }
}

export function formatRank(rank) {
    if(rank === null || rank === undefined) return "N/A";
    return rank;
}

/**
 * @param date {string}
 */
export function formatDate(date) {
    let d = new Date(date)

    if(isNaN(d.getDate())) {
      return "Unknown";
    }

    const formatter = new Intl.DateTimeFormat('en', { month: 'short' });
    return `${formatter.format(d)} ${d.getDate()}, ${d.getFullYear()}`
}

export function checkNull(value) {
    if(value === null || value === undefined) return "N/A";
    return value;
}

export function formatStatus(status) {
    switch(status) {
        case "finished_airing":
            return "<span style='color: #3F51B5;'>Finished Airing</span>";
        case "currently_airing":
            return "<span style='color: #4CAF50;'>Currently Airing</span>";
        case "not_yet_aired":
            return "<span style='color: #F44336;'>Not Yet Aired</span>";
        default:
            return status;
    }
}