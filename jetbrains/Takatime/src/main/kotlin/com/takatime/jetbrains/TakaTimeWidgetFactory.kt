package com.takatime.jetbrains

import com.intellij.icons.AllIcons
import com.intellij.openapi.project.Project
import com.intellij.openapi.ui.Messages
import com.intellij.openapi.ui.popup.JBPopupFactory
import com.intellij.openapi.ui.popup.ListPopup
import com.intellij.openapi.ui.popup.PopupStep
import com.intellij.openapi.ui.popup.util.BaseListPopupStep
import com.intellij.openapi.wm.StatusBar
import com.intellij.openapi.wm.StatusBarWidget
import com.intellij.openapi.wm.StatusBarWidgetFactory
import com.intellij.util.Consumer
import org.jetbrains.plugins.terminal.TerminalView
import java.awt.event.MouseEvent
import java.io.File

class TakaTimeWidgetFactory : StatusBarWidgetFactory {
    override fun getId() = "TakaTimeWidget"
    override fun getDisplayName() = "TakaTime Configuration"
    override fun isAvailable(project: Project) = true
    override fun createWidget(project: Project) = TakaTimeWidget(project)
    override fun disposeWidget(widget: StatusBarWidget) {}
    override fun canBeEnabledOn(statusBar: StatusBar) = true
}

class TakaTimeWidget(private val project: Project) : StatusBarWidget, StatusBarWidget.MultipleTextValuesPresentation {

    override fun ID() = "TakaTimeWidget"
    override fun getPresentation(): StatusBarWidget.WidgetPresentation = this
    override fun getSelectedValue() = "TakaTime ${TAKATIME_VERSION}"
    override fun getTooltipText() = "Click to open TakaTime Menu"

    // 👇 THE FIX: Wrap the blueprint in the JBPopupFactory!
    override fun getPopupStep(): ListPopup? {

        val options = listOf(
            "Open Dashboard",
            "Configure MongoDB URI"
        )

        val step = object : BaseListPopupStep<String>("TakaTime Menu", options) {

            override fun getIconFor(value: String) = when (value) {
                "Open Dashboard" -> AllIcons.General.Web
                "Configure MongoDB URI" -> AllIcons.General.Settings
                else -> null
            }

            override fun onChosen(
                selectedValue: String,
                finalChoice: Boolean
            ): PopupStep<*>? {

                return doFinalStep {
                    when (selectedValue) {
                        "Open Dashboard" -> launchDashboard()
                        "Configure MongoDB URI" -> configureUri()
                    }
                }
            }
        }

        return JBPopupFactory
            .getInstance()
            .createListPopup(step)
    }

    override fun getClickConsumer(): Consumer<MouseEvent>? = null

    private fun launchDashboard() {
        val mongoUri = TakaTimeConfig.getMongoUri()
        if (mongoUri.isNullOrBlank()) {
            Messages.showWarningDialog("Please configure your MongoDB URI first.", "TakaTime Setup Required")
            return
        }

        val homeDir = System.getProperty("user.home")
        val isWindows = System.getProperty("os.name").lowercase().contains("win")
        val ext = if (isWindows) ".exe" else ""
        val dashBinary = File(homeDir, ".takatime/bin/taka-dashboard-$TAKATIME_VERSION$ext")

        if (!dashBinary.exists()) {
            Messages.showWarningDialog("Dashboard binary not found. Please wait for the background update to finish.", "TakaTime Loading")
            return
        }

        try {
            val terminalView = TerminalView.getInstance(project)
            val widget = terminalView.createLocalShellWidget(project.basePath, "TakaTime Dashboard")
            val command = "\"${dashBinary.absolutePath}\" -MongoDBString \"$mongoUri\""
            widget.executeCommand(command)
        } catch (ex: Exception) {
            Messages.showErrorDialog("Failed to launch dashboard in terminal: ${ex.message}", "TakaTime Error")
        }
    }

    private fun configureUri() {
        val homeDir = System.getProperty("user.home")
        val configFile = File(homeDir, ".takatime.json")

        var savedUrl = ""
        var fileContent = ""

        if (configFile.exists()) {
            fileContent = configFile.readText()
            val match = """"MONGO_URI"\s*:\s*"([^"]+)"""".toRegex().find(fileContent)
            if (match != null) {
                savedUrl = match.groupValues[1]
            }
        }

        val newUrl = Messages.showInputDialog(
            project,
            "Enter your TakaTime MongoDB URI:",
            "TakaTime Setup",
            Messages.getQuestionIcon(),
            savedUrl,
            null
        )

        if (newUrl != null) {
            val jsonOutput = if (configFile.exists() && fileContent.contains("\"MONGO_URI\"")) {
                fileContent.replace(""""MONGO_URI"\s*:\s*"[^"]*"""".toRegex(), """"MONGO_URI": "$newUrl"""")
            } else {
                """{
    "MONGO_URI": "$newUrl",
    "VERSION": "$TAKATIME_VERSION"
}"""
            }

            configFile.writeText(jsonOutput)
            Messages.showInfoMessage(project, "MongoDB URL successfully saved!", "TakaTime")
            TakaTimeBinaryManager.checkAndDownloadIfNeeded(project)
        }
    }

    override fun install(statusBar: StatusBar) {}
    override fun dispose() {}
}