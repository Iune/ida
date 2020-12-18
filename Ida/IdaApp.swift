//
//  IdaApp.swift
//  Ida
//
//  Created by Aditya Duri on 12/17/20.
//

import SwiftUI

@main
struct IdaApp: App {
    var body: some Scene {
        WindowGroup {
            ContentView()
        }.commands {
            CommandGroup(after: CommandGroupPlacement.newItem) {
                Button("Open Contest JSON File") {
                    print("after item")
                }
            }
        }
    }
}
