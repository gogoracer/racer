// Trigger diffapply, should always be last
function diffApply() {
    document.querySelectorAll("[hon*=diffapply]").forEach(function (el) {
        const ids = hlive.getEventHandlerIDs(el);

        if (!ids["diffapply"]) {
            return;
        }

        // for (let i = 0; i < ids["diffapply"].length; i++) {
        hlive.sendMsg({
            type: 2,
            ids: ids["diffapply"],
            selected: false,
            valueMulti: [],
            data: {},
            extra: {},
        });
        // }
    });
}

// Register plugin
if (hlive.afterMessage.get("hdiffApply") === undefined) {
    hlive.afterMessage.set("hdiffApply", diffApply);
}
