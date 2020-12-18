//
//  ContentView.swift
//  Ida
//
//  Created by Aditya Duri on 12/17/20.
//

import SwiftUI
import IdaParser

struct ContentView: View {
    
    @State private var parser: Parser? = nil
    
    // Voter details
    @State private var voterName: String = ""
    @State private var voterMsg: String = ""
    @State private var parsedVotes: ParsedVotes? = nil
    
    var body: some View {
        VStack(alignment: .leading) {
            HStack {
                VStack(alignment: .leading) {
                    TextField("Voter Name", text: $voterName).disableAutocorrection(true)
                    TextEditor(text: $voterMsg).disableAutocorrection(true)
                        .font(.system(.body, design: .monospaced))
                }.frame(minWidth: 0, maxWidth: .infinity)
                Divider()
                VStack(alignment: .leading) {
                    ScrollView {
                        VStack(alignment: .leading) {
                            ForEach(self.parsedVotes?.votes ?? [], id: \.self) { vote in
                                Label(vote.entry.country.primaryName, systemImage: "\(vote.points).circle")
                                    .frame(alignment: .topLeading)
                            }
                        }
                    }.frame(minWidth: 0, maxWidth: .infinity)
                    if let parsedVotes = self.parsedVotes {
                        if parsedVotes.warningMsgs.count > 0 {
                            Divider()
                        }
                    }
                    VStack(alignment: .leading) {
                        ForEach(self.parsedVotes?.warningMsgs ?? [], id: \.self) { warning in
                            Label(warning, systemImage: "exclamationmark.triangle")
                                .font(.caption)
                        }
                    }
                }.frame(minWidth: 0, maxWidth: .infinity)
            }
        }
        .padding()
        .frame(minWidth: 500, maxWidth: .infinity, minHeight: 300, maxHeight: .infinity)
        .toolbar{
            ToolbarItem(placement: .primaryAction) {
                Button(action: selectContestFileAndSetParser) {
                    Label("Load Contest JSON File", systemImage: "doc")
                }.keyboardShortcut("O", modifiers: .command)
            }
            ToolbarItem(placement: .primaryAction) {
                Button(action: resetVoterDetails) {
                    Label("Reset Voter Details", systemImage: "arrow.clockwise")
                }.keyboardShortcut("R", modifiers: .command)
            }
            ToolbarItem(placement: .primaryAction) {
                Button(action: parseVotes) {
                    Label("Parse Votes", systemImage: "play.fill")
                }.disabled(self.parser == nil || self.voterName.isEmpty || self.voterMsg.isEmpty)
                .keyboardShortcut("C")
            }
        }
    }
    
    private func selectContestFileAndSetParser() {
        let dialog = NSOpenPanel()
        dialog.title = "Load contest JSON file"
        dialog.showsResizeIndicator = true
        dialog.showsHiddenFiles = false
        dialog.allowsMultipleSelection = false
        dialog.canChooseDirectories = false
        
        if (dialog.runModal() ==  NSApplication.ModalResponse.OK) {
            if let result = dialog.url {
                let path: String = result.path
                if let contest = try? Contest(atPath: path) {
                    self.parser = Parser(contest: contest)
                }
            }
        }
    }
    
    private func resetVoterDetails() {
        self.voterName = ""
        self.voterMsg = ""
        self.parsedVotes = Optional.none
    }
    
    private func parseVotes() {
        if let parser = self.parser {
            if !self.voterName.isEmpty && !self.voterMsg.isEmpty {
                let voter: Country? = parser.contest.findVoter(country: self.voterName)
                let lines: [String] = voterMsg.components(separatedBy: "\n")
                self.parsedVotes = parser.parse(voter: voter, lines: lines)
                self.copyVotesToClipboard(voter: voter, votes: self.parsedVotes!.votes)
            }
        }
    }
    
    private func copyVotesToClipboard(voter: Country?, votes: [Vote]) {
        if let contest = self.parser?.contest {
            var votesList = Array(repeating: "", count: contest.entries.count)
            if let voter = voter {
                if let index = contest.entries.map({$0.country}).firstIndex(of: voter) {
                    votesList[index] = "X"
                }
            }
            
            for vote in votes {
                if let index = contest.entries.firstIndex(of: vote.entry) {
                    votesList[index] = String(vote.points)
                }
            }
            
            let toCopy = votesList.joined(separator: "\n")
            let pasteBoard = NSPasteboard.general
            pasteBoard.clearContents()
            pasteBoard.setString(toCopy, forType: .string)
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
