// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.6;

/**
 * @author Antoine Falquerho
 * @title Forum example
 * @dev ...
 */
contract Forum {
    uint public last_post_id;
    mapping (address => User) private users;
    mapping (uint => Post) public posts;
    mapping(address => mapping(uint => Post)) public userPosts;
    mapping(uint => mapping(address => Vote)) public postVotes;

    struct User {
        string name;
        int karma;
        uint post_count;
    }

    struct Post {
        address owner;
        string title;
        string content;
        int karma;
        uint256 postedAt;
    }

    struct Vote {
        address owner;
        int vote;
    }

    function setName(address _owner, string memory _name) public{
        users[_owner].name = _name;
    }

    function getUser(address _address) public view returns (User memory) {
        return users[_address];
    }

    function addPost(address _owner, string memory title, string memory content) public {
        Post memory post = Post(_owner, title, content, 0, block.timestamp);
        posts[last_post_id] = post;
        uint user_last_post_id = users[_owner].post_count;
        userPosts[_owner][user_last_post_id] = post;
        last_post_id += 1;
        users[_owner].post_count = user_last_post_id + 1;
    }

    function addVote(address _owner, uint _post_id, int _vote) public {
        Vote memory vote = Vote(msg.sender, _vote);
        Post memory post = posts[_post_id];
        if(postVotes[_post_id][_owner].owner == address(0)){
            // Update Post karma
            posts[_post_id].karma += _vote;
            // Update User karma
            users[post.owner].karma += _vote;
        } else {
            Vote memory currentVote = postVotes[_post_id][_owner];
            if(currentVote.vote != _vote){
                if(_vote == 1){
                    _vote = 2;
                } else {
                    _vote = -2;
                }
                // Update Post karma
                posts[_post_id].karma += _vote;
                // Update User karma
                users[post.owner].karma += _vote;
            }
        }
        postVotes[_post_id][_owner] = vote;
    }
}