package com.takatime.jetbrains

import com.intellij.openapi.application.ApplicationManager
import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.progress.ProgressManager
import com.intellij.openapi.progress.Task
import com.intellij.openapi.project.Project
import com.intellij.openapi.ui.Messages
import java.io.File
import java.io.FileOutputStream
import java.net.HttpURLConnection
import java.net.URL

object TakaTimeBinaryManager {

    private fun binariesExist(targetVersion: String): Boolean {
        val homeDir = System.getProperty("user.home")
        val binDir = File(homeDir, ".takatime/bin")

        val isWindows = System.getProperty("os.name").lowercase().contains("win")
        val ext = if (isWindows) ".exe" else ""

        val uploadFile = File(binDir, "taka-upload-$targetVersion$ext")
        val dashFile = File(binDir, "taka-dashboard-$targetVersion$ext")

        return uploadFile.exists() && dashFile.exists()
    }

    private fun upgradeConfigVersionIfNeeded(): Boolean {
        val homeDir = System.getProperty("user.home")
        val configFile = File(homeDir, ".takatime.json")

        if (!configFile.exists()) return true

        val content = configFile.readText()
        val match = """"VERSION"\s*:\s*"([^"]+)"""".toRegex().find(content)
        val installedVersion = match?.groupValues?.get(1) ?: "v0.0.0"

        val installedNum = installedVersion.replace("v", "").replace(".", "").toIntOrNull() ?: 0
        val pluginNum = TAKATIME_VERSION.replace("v", "").replace(".", "").toIntOrNull() ?: 0

        if (installedNum >= pluginNum) return false

        println("TakaTime: Upgrading config from $installedVersion to $TAKATIME_VERSION")
        val newContent = if (content.contains("\"VERSION\"")) {
            content.replace(""""VERSION"\s*:\s*"[^"]*"""".toRegex(), """"VERSION": "$TAKATIME_VERSION"""")
        } else {
            content.trimEnd().removeSuffix("}") + ",\n    \"VERSION\": \"$TAKATIME_VERSION\"\n}"
        }

        configFile.writeText(newContent)
        return true
    }

    fun checkAndDownloadIfNeeded(project: Project?) {
        // 👇 GUARANTEES NO UI FREEZING: Run disk checks on a background thread
        ApplicationManager.getApplication().executeOnPooledThread {
            val versionChanged = upgradeConfigVersionIfNeeded()
            val missingBinaries = !binariesExist(TAKATIME_VERSION)

            if (versionChanged || missingBinaries) {
                // If we need to download, safely hop back to UI thread to launch the visual progress bar
                ApplicationManager.getApplication().invokeLater {
                    ensureBinaries(project, TAKATIME_VERSION)
                }
            }
        }
    }

    private fun getPlatformFilename(baseName: String): String? {
        val osName = System.getProperty("os.name").lowercase()
        val archName = System.getProperty("os.arch").lowercase()

        val osStr = when {
            osName.contains("win") -> "windows"
            osName.contains("mac") || osName.contains("darwin") -> "darwin"
            osName.contains("nix") || osName.contains("nux") || osName.contains("aix") -> "linux"
            else -> return null
        }

        val archStr = when {
            archName.contains("amd64") || archName.contains("x86_64") -> "amd64"
            archName.contains("aarch64") || archName.contains("arm64") -> "arm64"
            else -> return null
        }

        val ext = if (osStr == "windows") ".exe" else ""
        return "$baseName-$osStr-$archStr$ext"
    }

    private fun downloadSingleBinary(urlString: String, destFile: File) {
        var url = URL(urlString)
        var connection = url.openConnection() as HttpURLConnection
        var redirectCount = 0

        while (connection.responseCode in 300..399 && redirectCount < 5) {
            val newUrl = connection.getHeaderField("Location")
            url = URL(newUrl)
            connection = url.openConnection() as HttpURLConnection
            redirectCount++
        }

        if (connection.responseCode != 200) {
            throw Exception("HTTP ${connection.responseCode}")
        }

        connection.inputStream.use { input ->
            FileOutputStream(destFile).use { output ->
                input.copyTo(output)
            }
        }

        if (!System.getProperty("os.name").lowercase().contains("win")) {
            destFile.setExecutable(true, false)
        }
    }

    private fun ensureBinaries(project: Project?, version: String) {
        val requiredBinaries = listOf("taka-upload", "taka-dashboard")
        val homeDir = System.getProperty("user.home")
        val binDir = File(homeDir, ".takatime/bin")

        if (!binDir.exists()) {
            binDir.mkdirs()
        }

        ProgressManager.getInstance().run(object : Task.Backgroundable(project, "Installing TakaTime $version...", false) {
            override fun run(indicator: ProgressIndicator) {
                try {
                    for (baseName in requiredBinaries) {
                        val filename = getPlatformFilename(baseName)
                            ?: throw Exception("Unsupported Platform: ${System.getProperty("os.name")} - ${System.getProperty("os.arch")}")

                        val ext = if (System.getProperty("os.name").lowercase().contains("win")) ".exe" else ""
                        val localFilename = "$baseName-$version$ext"
                        val destFile = File(binDir, localFilename)
                        val url = "https://github.com/Rtarun3606k/TakaTime/releases/download/$version/$filename"

                        indicator.text = "Fetching $baseName..."

                        if (destFile.exists()) {
                            destFile.delete()
                        }

                        downloadSingleBinary(url, destFile)
                    }

                    ApplicationManager.getApplication().invokeLater {
                        Messages.showInfoMessage("TakaTime updated to $version successfully!", "TakaTime")
                    }

                } catch (e: Exception) {
                    ApplicationManager.getApplication().invokeLater {
                        Messages.showErrorDialog("TakaTime Update Failed: ${e.message}", "TakaTime Error")
                    }
                }
            }
        })
    }
}