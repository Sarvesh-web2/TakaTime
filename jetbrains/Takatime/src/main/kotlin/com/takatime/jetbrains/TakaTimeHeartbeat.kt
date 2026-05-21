package com.takatime.jetbrains

import com.intellij.openapi.editor.Document

object TakaTimeHeartbeat {

    private var lastHeartbeatTime: Long = 0
    private const val COOLDOWN_MS: Long = 120_000

    fun handleHeartbeat(document: Document) {
        val now = System.currentTimeMillis()

        // 1. Check Cooldown (blocks spam)
        if (lastHeartbeatTime != 0L && (now - lastHeartbeatTime < COOLDOWN_MS)) return

        // 2. Calculate true duration in seconds
        var durationSeconds = (now - lastHeartbeatTime) / 1000

        // 3. The AFK Guard: If this is their first keystroke ever, OR they
        // took a break longer than your cooldown, reset the duration to 120.
        if (lastHeartbeatTime == 0L || durationSeconds > 120) {
            durationSeconds = 120
        }

        // 4. Pass the dynamic duration to the uploader!
        if (!TakaTimeUploader.spawnProcess(document, durationSeconds.toString())) return

        // 5. Reset Timer
        lastHeartbeatTime = now
    }
}